use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct FileData {
    pub id: u64,
    pub name: String,
    pub is_dir: bool,
    pub size: u64,
    pub blocks: Option<Vec<u64>>,
    pub timestamp: String,
}
