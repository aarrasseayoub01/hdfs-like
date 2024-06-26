package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	utils "github.com/aarrasseayoub01/namenode/namenode/internal/fs"
	"github.com/aarrasseayoub01/namenode/namenode/internal/gRPC"
	"github.com/aarrasseayoub01/namenode/namenode/internal/persistence"
)

type FileSystemService struct {
	rootDirectory *utils.Directory
	rootMutex     sync.RWMutex
}

func NewFileSystemService(root *utils.Directory) *FileSystemService {
	return &FileSystemService{rootDirectory: root}
}

// CreateFile creates a new file in the file system.
// func (fs *FileSystemService) CreateFile(filePath string) (*utils.Inode, error) {
// 	fs.rootMutex.Lock()
// 	defer fs.rootMutex.Unlock()
// 	dirPath, fileName := filepath.Split(filePath)
// 	parentDir := utils.FindDirectory(fs.rootDirectory, dirPath)
// 	if parentDir == nil {
// 		return nil, fmt.Errorf("parent directory does not exist")
// 	}
// 	if _, exists := parentDir.ChildFiles[fileName]; exists {
// 		return nil, fmt.Errorf("file already exists")
// 	}

// 	newFileInode := &utils.Inode{
// 		ID:        utils.GenerateInodeID(),
// 		Name:      fileName,
// 		IsDir:     false,
// 		Size:      0,
// 		Timestamp: time.Now(),
// 	}
// 	parentDir.ChildFiles[fileName] = newFileInode
// 	persistence.RecordEditLog("CREATE_FILE", filePath, newFileInode, blocks)

// 	return newFileInode, nil
// }

// ReadFile reads a file in the file system.
func (fs *FileSystemService) ReadFile(filePath string) (*utils.Inode, error) {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()
	dirPath, fileName := filepath.Split(filePath)
	parentDir := utils.FindDirectory(fs.rootDirectory, dirPath)
	if parentDir == nil {
		return nil, fmt.Errorf("parent directory does not exist")
	}
	if _, exists := parentDir.ChildFiles[fileName]; !exists {
		return nil, fmt.Errorf("file doesn't exists")
	}

	return parentDir.ChildFiles[fileName], nil
}

// DeleteFile deletes a file from the file system.
func (fs *FileSystemService) DeleteFile(filePath string) error {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()

	dirPath, fileName := filepath.Split(filePath)
	parentDir := utils.FindDirectory(fs.rootDirectory, dirPath)
	if parentDir == nil {
		return fmt.Errorf("directory does not exist")
	}
	if _, exists := parentDir.ChildFiles[fileName]; !exists {
		return fmt.Errorf("file does not exist")
	}

	delete(parentDir.ChildFiles, fileName)

	persistence.RecordEditLog("DELETE_FILE", filePath, nil)

	return nil
}

// CreateDirectory creates a new directory in the file system.
func (fs *FileSystemService) CreateDirectory(dirPath string) (*utils.Inode, error) {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()
	parentPath, dirName := filepath.Dir(dirPath), filepath.Base(dirPath)
	parentDir := utils.FindDirectory(fs.rootDirectory, parentPath)
	if parentDir == nil {
		return nil, fmt.Errorf("No parent path provided " + parentPath)
	}
	if _, exists := parentDir.ChildDirs[dirName]; exists {
		return nil, fmt.Errorf("directory already exists")
	}

	newDirInode := &utils.Inode{
		ID:        utils.GenerateInodeID(),
		Name:      dirName,
		IsDir:     true,
		Blocks:    []utils.BlockAssignment{},
		Timestamp: time.Now(),
	}
	parentDir.ChildDirs[dirName] = &utils.Directory{
		Inode:      newDirInode,
		ChildFiles: make(map[string]*utils.Inode),
		ChildDirs:  make(map[string]*utils.Directory),
	}

	persistence.RecordEditLog("CREATE_DIRECTORY", dirPath, newDirInode)

	return newDirInode, nil
}

func (fs *FileSystemService) ReadDirectory(dirPath string) ([]*utils.Inode, error) {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()

	// parentPath, dirName := filepath.Dir(dirPath), filepath.Base(dirPath)

	// Find the parent directory
	dir := utils.FindDirectory(fs.rootDirectory, dirPath)
	fmt.Println(dirPath)
	if dir == nil {
		return nil, fmt.Errorf("no parent path provided: %s", dirPath)
	}

	// // Check if the directory exists
	// dir, exists := parentDir.ChildDirs[dirName]
	// if !exists {
	// 	return nil, fmt.Errorf("directory does not exist: %s", dirPath)
	// }

	// Read child files and directories
	childFiles := make([]*utils.Inode, 0, len(dir.ChildFiles))
	fmt.Println(dirPath)
	for _, inode := range dir.ChildFiles {
		childFiles = append(childFiles, inode)
	}

	childDirs := make([]*utils.Inode, 0, len(dir.ChildDirs))
	for _, dir := range dir.ChildDirs {
		childDirs = append(childDirs, dir.Inode)
	}
	childFiles = append(childFiles, childDirs...)

	return childFiles, nil
}

// DeleteDirectory deletes a directory from the file system.
func (fs *FileSystemService) DeleteDirectory(dirPath string) error {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()

	parentPath, dirName := filepath.Dir(dirPath), filepath.Base(dirPath)
	parentDir := utils.FindDirectory(fs.rootDirectory, parentPath)
	if parentDir == nil {
		return fmt.Errorf("directory does not exist")
	}
	if dir, exists := parentDir.ChildDirs[dirName]; !exists || len(dir.ChildFiles) > 0 || len(dir.ChildDirs) > 0 {
		return fmt.Errorf("directory is not empty or does not exist")
	}

	delete(parentDir.ChildDirs, dirName)

	persistence.RecordEditLog("DELETE_DIRECTORY", dirPath, nil)

	return nil
}

func (fs *FileSystemService) CreateFile(filePath string, fileSize int64) (*utils.Inode, error) {
	const blockSize int64 = 64 * 1024 * 1024
	var blockAssignments []utils.BlockAssignment

	dataNodeManager := gRPC.GetInstance()
	dataNodes := dataNodeManager.GetDataNodes()
	// Calculate the number of blocks needed
	numBlocks := fileSize / blockSize
	if fileSize%blockSize != 0 {
		numBlocks++
	}

	// Prepare a slice of DataNode addresses for round-robin allocation
	dataNodeAddresses := make([]string, 0, len(dataNodes))

	if len(dataNodes) == 0 {
		return nil, errors.New("There are no DataNodes")
	}
	for address := range dataNodes {
		dataNodeAddresses = append(dataNodeAddresses, address)
	}

	// Assign blocks to DataNodes
	for i := int64(0); i < numBlocks; i++ {
		blockID := fmt.Sprintf("%s-block-%d", filepath.Base(filePath), i)
		dataNodeIndex := i % int64(len(dataNodeAddresses)) // Round-robin allocation

		blockAssignments = append(blockAssignments, utils.BlockAssignment{
			BlockID:           blockID,
			DataNodeAddresses: []string{dataNodeAddresses[dataNodeIndex]},
		})
	}

	dirPath, fileName := filepath.Split(filePath)
	parentDir := utils.FindDirectory(fs.rootDirectory, dirPath)
	if parentDir == nil {
		return nil, fmt.Errorf("parent directory does not exist")
	}
	if _, exists := parentDir.ChildFiles[fileName]; exists {
		return nil, fmt.Errorf("file already exists")
	}

	newFileInode := &utils.Inode{
		ID:        utils.GenerateInodeID(),
		Name:      fileName,
		IsDir:     false,
		Size:      fileSize,
		Blocks:    utils.AllocateFileBlocksResponse{BlockAssignments: blockAssignments}.BlockAssignments,
		Timestamp: time.Now(),
	}
	parentDir.ChildFiles[fileName] = newFileInode
	persistence.RecordEditLog("CREATE_FILE", filePath, newFileInode)

	// return &utils.AllocateFileBlocksResponse{BlockAssignments: blockAssignments}, nil
	return newFileInode, nil

}
