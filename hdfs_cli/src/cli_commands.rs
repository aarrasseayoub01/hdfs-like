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

pub async fn handle_create_subcommand(
    create_matches: &ArgMatches,
    client: &Client,
    base_url: &str,
) -> Result<(), reqwest::Error> {
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
    match client.post(&url).json(&post_body).send().await {
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
