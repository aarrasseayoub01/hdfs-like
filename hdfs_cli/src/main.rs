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
        let local_path = create_matches
            .value_of("local_path")
            .expect("Required argument missing: local_path");

        let server_path = create_matches
            .value_of("server_path")
            .expect("Required argument missing: server_path");
        let is_directory = create_matches.is_present("directory");

        let _ = cli_commands::handle_create_subcommand(
            local_path,
            server_path,
            &client,
            base_url,
            is_directory,
        )
        .await;
    } else if let Some(("read", read_matches)) = matches.subcommand() {
        let _ = cli_commands::handle_read_subcommand(read_matches, &client, base_url).await;
    } else {
        println!("Invalid command");
    }
}
