use clap::{App, Arg, SubCommand};
use prettytable::{format, row, Table};
use reqwest::Client;
use serde::{Deserialize, Serialize};
use serde_json;
use serde_json::json;
use tokio;

#[tokio::main]
async fn main() {
    let matches = build_cli_app().get_matches();

    let client = Client::new();
    let base_url = "http://localhost:8080";

    if let Some(("create", create_matches)) = matches.subcommand() {
        handle_create_subcommand(create_matches, &client, base_url).await;
    } else if let Some(("read", read_matches)) = matches.subcommand() {
        handle_read_subcommand(read_matches, &client, base_url).await;
    } else {
        println!("Invalid command");
    }
}

fn build_cli_app() -> App<'static> {
    App::new("Rust CLI for REST Requests")
        .version("1.0")
        .author("Your Name")
        .about("Performs HTTP requests")
        .subcommand(
            SubCommand::with_name("create")
                .about("Creates a file or directory with the specified path")
                .arg(Arg::with_name("path").required(true).index(1))
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

async fn handle_create_subcommand(
    create_matches: &clap::ArgMatches,
    client: &Client,
    base_url: &str,
) {
    let path = create_matches
        .value_of("path")
        .expect("Required argument missing: path");

    let (endpoint, key) = if create_matches.is_present("directory") {
        ("createDir", "dirPath")
    } else {
        ("createFile", "filePath")
    };

    let post_body = json!({ key: path });

    let url = format!("{}/{}", base_url, endpoint);
    let response = client
        .post(&url)
        .json(&post_body)
        .send()
        .await
        .expect("Failed to send request");
    println!(
        "Response: {:?}",
        response.text().await.expect("Failed to read response")
    );
}

// Define a struct matching your JSON structure
#[derive(Debug, Serialize, Deserialize)]
struct FileData {
    #[serde(rename = "ID")]
    id: u64,
    #[serde(rename = "Name")]
    name: String,
    #[serde(rename = "IsDir")]
    is_dir: bool,
    #[serde(rename = "Size")]
    size: u64,
    #[serde(rename = "Blocks")]
    blocks: Option<Vec<u64>>, // adjust the type according to your actual data
    #[serde(rename = "Timestamp")]
    timestamp: String,
}
async fn handle_read_subcommand(read_matches: &clap::ArgMatches, client: &Client, base_url: &str) {
    let path = read_matches
        .value_of("path")
        .expect("Required argument missing: path");

    let (endpoint, key) = if read_matches.is_present("directory") {
        ("readDir", "dirPath")
    } else {
        ("readFile", "filePath")
    };

    let url = format!("{}/{}?{}={}", base_url, endpoint, key, path);
    let response = client
        .get(&url)
        .send()
        .await
        .expect("Failed to send request");

    let response_text = response.text().await.expect("Failed to read response");

    if read_matches.is_present("directory") {
        let files: Vec<FileData> = serde_json::from_str(&response_text)
            .expect("Failed to parse JSON as a list of FileData");

        // Use the new method to print the file data
        print_file_data_table(&files);
    } else {
        // Parse the response as a single FileData
        let file_data: FileData =
            serde_json::from_str(&response_text).expect("Failed to parse JSON as FileData");

        // Print the file data
        println!("Parsed Response: {:?}", file_data);
    }
}

fn print_file_data_table(files: &[FileData]) {
    let mut table = Table::new();
    // Set the format with separate lines for each row and column
    // table.set_format(*format::consts::FORMAT_BORDERS_ONLY);

    // Add a header row
    table.add_row(row!["ID", "Name", "Type", "Size", "Blocks", "Timestamp"]);

    for file in files {
        // Determine file type
        let file_type = if file.is_dir { "Dir" } else { "File" };

        // Display blocks if available
        let blocks_display = match &file.blocks {
            Some(blocks) => format!("{:?}", blocks),
            None => "None".to_string(),
        };

        // Add each row to the table
        table.add_row(row![
            file.id,
            file.name,
            file_type,
            file.size,
            blocks_display,
            file.timestamp
        ]);
    }

    // Print the table to the console
    table.printstd();
}
