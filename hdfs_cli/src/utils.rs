use crate::models::BlockAssignment;
use base64;

use anyhow::{Context, Result};
use std::env;
use std::fs::File;
use std::path::Path;

use std::io::{Read, Seek, SeekFrom};
const BLOCK_SIZE: u64 = 64 * 1024 * 1024; // 64 MB

pub async fn process_file_in_blocks(
    local_path: &str,
    block_assignments: &Vec<BlockAssignment>,
) -> Result<()> {
    let current = &get_current_working_dir();
    let current_dir = Path::new(current);
    let mut file = File::open(current_dir.join(local_path)).context("Failed to open file")?;

    let file_size = file
        .metadata()
        .context("Failed to read file metadata")?
        .len();
    let num_blocks = (file_size + BLOCK_SIZE - 1) / BLOCK_SIZE;

    for (block_num, block_assignment) in block_assignments
        .iter()
        .enumerate()
        .take(num_blocks as usize)
    {
        let mut buffer = Vec::new();
        file.seek(SeekFrom::Start(block_num as u64 * BLOCK_SIZE))
            .context("Failed to seek in file")?;

        let block_size = if block_num as u64 == num_blocks - 1 {
            file_size - block_num as u64 * BLOCK_SIZE
        } else {
            BLOCK_SIZE
        };

        buffer.resize(block_size as usize, 0);
        file.read_exact(&mut buffer)
            .context("Failed to read file block")?;

        if let Some(datanode_address) = block_assignment.datanode_addresses.get(0) {
            send_block_to_datanode(&block_assignment.block_id, datanode_address, &buffer)
                .await
                .with_context(|| {
                    format!("Failed to send block to datanode {}", datanode_address)
                })?;
        }
    }

    Ok(())
}
async fn send_block_to_datanode(
    block_id: &str,
    _datanode_address: &str,
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
    _datanode_address: &str,
) -> Result<Vec<u8>, Box<dyn Error>> {
    let url = format!("http://localhost:8081/getBlock/{}", block_id);
    print!("{}", url);
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
