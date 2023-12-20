use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct FileData {
    #[serde(rename = "ID")]
    pub id: u64,
    #[serde(rename = "Name")]
    pub name: String,
    #[serde(rename = "IsDir")]
    pub is_dir: bool,
    #[serde(rename = "Size")]
    pub size: u64,
    #[serde(rename = "Blocks")]
    pub blocks: Option<Vec<u64>>,
    #[serde(rename = "Timestamp")]
    pub timestamp: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct BlockAssignment {
    #[serde(rename = "blockId")]
    pub block_id: String,
    #[serde(rename = "datanodeAddresses")]
    pub datanode_addresses: Vec<String>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct AllocationResponse {
    #[serde(rename = "blockAssignments")]
    pub block_assignments: Vec<BlockAssignment>,
}
