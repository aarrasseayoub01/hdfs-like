use clap::{App, Arg, ArgMatches, SubCommand};
use prettytable::{color, row, Attr, Cell, Row, Table};
use reqwest::Client;
use serde_json::json;

use crate::models::AllocationResponse;
// use crate::models::BlockAssignment;
use crate::models::FileData;
use crate::utils::process_file_in_blocks;
use crate::utils::retrieve_block_from_datanode;

pub fn build_cli_app() -> App<'static> {
    App::new("Rust CLI for REST Requests")
        .version("1.0")
        .author("Ayoub")
        .about("Performs HTTP requests")
        .subcommand(
            SubCommand::with_name("create")
                .about("Creates a file or directory with the specified path")
                .arg(
                    Arg::with_name("local_path")
                        .help("The local path of the file")
                        .required(true)
                        .index(1),
                ) // First argument for 'create'
                .arg(
                    Arg::with_name("server_path")
                        .help("The server path of the file")
                        .required(true)
                        .index(2),
                ) // Second argument for 'create'
                .arg(
                    Arg::with_name("directory")
                        .short('d')
                        .long("directory")
                        .help("Create a directory instead of a file"),
                ),
        )
        .subcommand(
            SubCommand::with_name("read")
                .about("Reads a file or directory with the specified path")
                .arg(Arg::with_name("path").required(true).index(1))
                .arg(
                    Arg::with_name("directory")
                        .short('d')
                        .long("directory")
                        .help("Read a directory instead of a file"),
                ),
        )
        .subcommand(
            SubCommand::with_name("get")
                .about("Retrieves a file from the data nodes")
                .arg(
                    Arg::with_name("filepath")
                        .help("The path of the file to retrieve")
                        .required(true)
                        .index(1),
                ),
        )
}

pub async fn handle_create_subcommand(
    local_path: &str,
    server_path: &str,
    client: &Client,
    base_url: &str,
    is_directory: bool, // New parameter to indicate directory option
) -> Result<(), reqwest::Error> {
    // Determine the endpoint and request body key based on is_directory
    let (endpoint, key) = if is_directory {
        ("createDir", "dirPath")
    } else {
        ("createFile", "filePath")
    };

    // Step 1: Create the file's or directory's metadata
    let post_body = json!({ key: server_path });
    let create_url = format!("{}/{}", base_url, endpoint);
    let create_response = match client.post(&create_url).json(&post_body).send().await {
        Ok(res) => res,
        Err(e) => {
            if e.is_timeout() {
                println!("Request timed out. Please try again.");
            } else {
                println!("An error occurred: {}", e);
            }
            return Err(e);
        }
    };

    // Print the JSON response from the endpoint call
    let create_response_text = create_response
        .text()
        .await
        .expect("Failed to read response");
    println!(
        "Response from {} endpoint: {}",
        endpoint, create_response_text
    );
    if !is_directory {
        let allocate_url = format!("{}/allocate", base_url);
        let allocate_response = match client
            .post(&allocate_url)
            .json(&json!({"filePath": server_path, "fileSize": 40 })) // Use the server_path here as well
            .send()
            .await
        {
            Ok(res) => res,
            Err(e) => {
                if e.is_timeout() {
                    println!("Request timed out. Please try again.");
                } else {
                    println!("An error occurred: {}", e);
                }
                return Err(e);
            }
        };

        // Decode the JSON response
        let allocation_data: AllocationResponse = match allocate_response.json().await {
            Ok(data) => data,
            Err(e) => {
                println!("Failed to decode response: {}", e);
                return Err(e);
            }
        };

        process_file_in_blocks(local_path, &allocation_data.block_assignments).await;
    }
    Ok(())
}

pub async fn handle_read_subcommand(
    read_matches: &ArgMatches,
    client: &Client,
    base_url: &str,
) -> Result<(), reqwest::Error> {
    let path = read_matches
        .value_of("path")
        .expect("Required argument missing: path");

    let (endpoint, key) = if read_matches.is_present("directory") {
        ("readDir", "dirPath")
    } else {
        ("readFile", "filePath")
    };

    let url = format!("{}/{}?{}={}", base_url, endpoint, key, path);
    let response = match client.get(&url).send().await {
        Ok(res) => res,
        Err(e) => {
            if e.is_timeout() {
                println!("Request timed out. Please try again.");
            } else {
                println!("An error occurred: {}", e);
            }
            return Err(e);
        }
    };

    let response_text = response.text().await.expect("Failed to read response");
    if read_matches.is_present("directory") {
        let files: Vec<FileData> = serde_json::from_str(&response_text)
            .expect("Failed to parse JSON as a list of FileData");

        print_file_data_table(&files);
    } else {
        let file_data: FileData =
            serde_json::from_str(&response_text).expect("Failed to parse JSON as FileData");

        println!("Parsed Response: {:?}", file_data);
    }
    Ok(())
}

fn print_file_data_table(files: &[FileData]) {
    let mut table = Table::new();
    table.add_row(row!["ID", "Name", "Type", "Size", "Blocks", "Timestamp"]);

    for file in files {
        let file_type = if file.is_dir { "Dir" } else { "File" };
        let blocks_display = file
            .blocks
            .as_ref()
            .map_or_else(|| "None".to_string(), |b| format!("{:?}", b));
        let name_cell = if file.is_dir {
            Cell::new(&file.name).with_style(Attr::ForegroundColor(color::BLUE))
        } else {
            Cell::new(&file.name)
        };

        table.add_row(Row::new(vec![
            Cell::new(&file.id.to_string()),
            name_cell,
            Cell::new(file_type),
            Cell::new(&file.size.to_string()),
            Cell::new(&blocks_display),
            Cell::new(&file.timestamp),
        ]));
    }

    table.printstd();
}

// In your cli_commands module

pub async fn handle_get_subcommand(
    filepath: &str,
    client: &Client,
    base_url: &str,
) -> Result<(), reqwest::Error> {
    // Step 1: Call the name node to get block assignments and data node addresses
    // Assuming the endpoint is something like "/getFileBlocks?filePath={filepath}"
    let url = format!("{}/getFileBlocks?filePath={}", base_url, filepath);
    let response = client.get(&url).send().await?;
    let allocation_data: AllocationResponse = response.json().await?;

    // Step 2: Retrieve each block from its respective data node
    let mut file_data = Vec::new();
    for block_assignment in allocation_data.block_assignments {
        if let Some(datanode_address) = block_assignment.datanode_addresses.get(0) {
            match retrieve_block_from_datanode(&block_assignment.block_id, datanode_address).await {
                Ok(block_data) => {
                    println!("Retrieved block data: {:?}", block_data);
                    file_data.extend(block_data);
                }
                Err(e) => {
                    println!("Error retrieving block: {}", e);
                    continue;
                }
            }
        }
    }

    // Step 3: Write the gathered data to a file
    std::fs::write(filepath, file_data);

    Ok(())
}
