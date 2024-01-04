use crate::models::BlockAssignment;
use base64;
use core::num;
use std::env;
use std::fs::File;
use std::path::Path;
use std::path::PathBuf;

use std::io::{self, Read, Seek, SeekFrom};
const BLOCK_SIZE: u64 = 64 * 1024 * 1024; // 64 MB

pub async fn process_file_in_blocks(
    local_path: &str,
    block_assignments: &Vec<BlockAssignment>,
) -> io::Result<()> {
    let mut file = File::open(Path::new(&get_current_working_dir()).join(local_path))?;

    let file_size = file.metadata()?.len();
    let num_blocks = (file_size + BLOCK_SIZE - 1) / BLOCK_SIZE;

    for (block_num, block_assignment) in block_assignments
        .iter()
        .enumerate()
        .take(num_blocks as usize)
    {
        let mut buffer = Vec::new();
        file.seek(SeekFrom::Start(block_num as u64 * BLOCK_SIZE))?;

        let block_size = if block_num as u64 == num_blocks - 1 {
            file_size - block_num as u64 * BLOCK_SIZE
        } else {
            BLOCK_SIZE
        };

        buffer.resize(block_size as usize, 0);
        file.read_exact(&mut buffer)?;

        // Send the block to the data node
        if let Some(datanode_address) = block_assignment.datanode_addresses.get(0) {
            send_block_to_datanode(&block_assignment.block_id, datanode_address, &buffer).await;
            // println!("{}", datanode_address)
        }
    }

    Ok(())
}

async fn send_block_to_datanode(
    block_id: &str,
    datanode_address: &str,
    data: &[u8],
) -> Result<(), reqwest::Error> {
    let client = reqwest::Client::new();
    let url = format!("http://localhost:8081/addBlock");

    let block_request = serde_json::json!({
        "blockId": block_id,
        "data": base64::encode(data),
    });

    client
        .post(&url)
        .json(&block_request)
        .send()
        .await?
        .error_for_status()?;

    Ok(())
}

use reqwest;
use std::error::Error;

pub async fn retrieve_block_from_datanode(
    block_id: &str,
    datanode_address: &str,
) -> Result<Vec<u8>, Box<dyn Error>> {
    let url = format!("http://localhost:8081/getBlock/{}", block_id);

    let response = reqwest::get(&url).await?;
    let status = response.status();

    if !status.is_success() {
        return Err(format!("Failed to retrieve block: HTTP {}", status).into());
    }

    let data = response.bytes().await?.to_vec();
    Ok(data)
}

fn get_current_working_dir() -> String {
    let res = env::current_dir();
    match res {
        Ok(path) => path.into_os_string().into_string().unwrap(),
        Err(_) => "FAILED".to_string(),
    }
}
