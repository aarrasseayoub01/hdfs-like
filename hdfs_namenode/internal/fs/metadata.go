package fs

import "time"

type Inode struct {
	ID        int64
	Name      string
	IsDir     bool
	Size      int64
	Blocks    []int64
	Timestamp time.Time
}

type Directory struct {
	Inode      *Inode
	ChildFiles map[string]*Inode
	ChildDirs  map[string]*Directory
}

type File struct {
	Inode Inode
	// File-specific data
}

type Block struct {
	ID     int64
	Offset int64
	Size   int64
}
