use clap::{App, Arg, SubCommand};
use reqwest::Client;
use serde_json::json;
use tokio;

#[tokio::main]
async fn main() {
    let matches = build_cli_app().get_matches();

    let client = Client::new();
    let base_url = "http://localhost:8080";

    if let Some(("create", create_matches)) = matches.subcommand() {
        handle_create_subcommand(create_matches, &client, base_url).await;
    } else {
        println!("Invalid command");
    }
}

fn build_cli_app() -> App<'static> {
    App::new("Rust CLI for REST Requests")
        .version("1.0")
        .author("Your Name")
        .about("Performs HTTP requests")
        .subcommand(SubCommand::with_name("create")
            .about("Creates a file or directory with the specified path")
            .arg(Arg::with_name("path")
                .required(true)
                .index(1))
            .arg(Arg::with_name("directory")
                .short('d')
                .long("directory")
                .help("Create a directory instead of a file")))
}

async fn handle_create_subcommand(create_matches: &clap::ArgMatches, client: &Client, base_url: &str) {
    let path = create_matches.value_of("path").expect("Required argument missing: path");

    let (endpoint, key) = if create_matches.is_present("directory") {
        ("createDir", "dirPath")
    } else {
        ("createFile", "filePath")
    };

    let post_body = json!({ key: path });

    let url = format!("{}/{}", base_url, endpoint);
    let response = client.post(&url).json(&post_body).send().await.expect("Failed to send request");
    println!("Response: {:?}", response.text().await.expect("Failed to read response"));
}
