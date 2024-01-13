package fs

import "time"

type Inode struct {
	ID        int64
	Name      string
	IsDir     bool
	Size      int64
	Blocks    []BlockAssignment
	Timestamp time.Time
}

type Directory struct {
	Inode      *Inode
	ChildFiles map[string]*Inode
	ChildDirs  map[string]*Directory
}

type File struct {
	Inode Inode
}

type Block struct {
	ID     int64
	Offset int64
	Size   int64
}

type BlockAssignment struct {
	BlockID           string   `json:"blockId"`
	DataNodeAddresses []string `json:"datanodeAddresses"`
}

type AllocateFileBlocksResponse struct {
	BlockAssignments []BlockAssignment `json:"blockAssignments"`
}

type AllocateFileBlocksRequest struct {
	FilePath string `json:"filePath"`
	FileSize int64  `json:"fileSize"`
}
