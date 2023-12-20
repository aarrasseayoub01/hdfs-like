use crate::models::BlockAssignment;
use std::fs::File;
use std::io::{self, Read, Seek, SeekFrom};
const BLOCK_SIZE: u64 = 64 * 1024 * 1024; // 64 MB

pub async fn process_file_in_blocks(
    local_path: &str,
    block_assignments: &[BlockAssignment],
) -> io::Result<()> {
    let mut file = File::open(local_path)?;

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
            send_block_to_datanode(&block_assignment.block_id, datanode_address, &buffer).await?;
        }
    }

    Ok(())
}
