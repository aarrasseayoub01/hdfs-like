use clap::{App, Arg, SubCommand};
use reqwest::Client;
use tokio;

#[tokio::main]
async fn main() {
    let matches = App::new("Rust CLI for REST Requests")
        .version("1.0")
        .author("Your Name")
        .about("Performs HTTP requests")
        .subcommand(SubCommand::with_name("get")
            .about("performs a GET request")
            .arg(Arg::with_name("url").required(true)))
        .subcommand(SubCommand::with_name("post")
            .about("performs a POST request")
            .arg(Arg::with_name("url").required(true))
            .arg(Arg::with_name("body").required(true)))
        // Add similar subcommands for DELETE and PUT
        .subcommand(SubCommand::with_name("delete")
            .about("performs a DELETE request")
            .arg(Arg::with_name("url").required(true)))
        .subcommand(SubCommand::with_name("put")
            .about("performs a PUT request")
            .arg(Arg::with_name("url").required(true))
            .arg(Arg::with_name("body").required(true)))
        .get_matches();

    let client = Client::new();

    match matches.subcommand() {
        Some(("get", get_matches)) => {
            let url = get_matches.value_of("url").unwrap();
            let response = client.get(url).send().await.unwrap();
            println!("Response: {:?}", response.text().await.unwrap());
        }
        Some(("post", post_matches)) => {
            let url = post_matches.value_of("url").unwrap();
            let body = post_matches.value_of("body").unwrap();
            let response = client.post(url).body(body.to_string()).send().await.unwrap();
            println!("Response: {:?}", response.text().await.unwrap());
        }
        Some(("delete", delete_matches)) => {
            let url = delete_matches.value_of("url").unwrap();
            let response = client.delete(url).send().await.unwrap();
            println!("Response: {:?}", response.text().await.unwrap());
        }
        Some(("put", put_matches)) => {
            let url = put_matches.value_of("url").unwrap();
            let body = put_matches.value_of("body").unwrap();
            let response = client.put(url).body(body.to_string()).send().await.unwrap();
            println!("Response: {:?}", response.text().await.unwrap());
        }
        _ => {
            println!("Invalid command");
        }
    }
}
