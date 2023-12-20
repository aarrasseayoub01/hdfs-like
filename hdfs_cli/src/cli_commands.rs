use clap::{App, Arg, ArgMatches, SubCommand};
use prettytable::{color, row, Attr, Cell, Row, Table};
use reqwest::Client;
use serde_json::json;

use crate::models::FileData;

pub fn build_cli_app() -> App<'static> {
    App::new("Rust CLI for REST Requests")
        .version("1.0")
        .author("Your Name")
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

    // Extract the file ID from the response, assuming it's returned as JSON
    // let create_response_json: serde_json::Value = create_response.json().await?;
    // let file_id = create_response_json["fileId"].as_str().unwrap_or_default(); // Adjust the key as needed

    // Step 2: Call the allocate endpoint
    let allocate_url = format!("{}/allocateFileBlocks", base_url);
    match client
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

    // Handle the allocate response as needed
    // You can print the response or parse it depending on your requirements

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
