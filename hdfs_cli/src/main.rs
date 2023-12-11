use tokio;

mod cli_commands;
mod http_client;
mod models;

#[tokio::main]
async fn main() {
    let matches = cli_commands::build_cli_app().get_matches();

    let client = http_client::build_client();
    let base_url = "http://localhost:8080";

    if let Some(("create", create_matches)) = matches.subcommand() {
        cli_commands::handle_create_subcommand(create_matches, &client, base_url).await;
    } else if let Some(("read", read_matches)) = matches.subcommand() {
        cli_commands::handle_read_subcommand(read_matches, &client, base_url).await;
    } else {
        println!("Invalid command");
    }
}
