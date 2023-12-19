package service

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	utils "github.com/aarrasseayoub01/namenode/namenode/internal/fs"
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
func (fs *FileSystemService) CreateFile(filePath string) error {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()
	dirPath, fileName := filepath.Split(filePath)
	parentDir := utils.FindDirectory(fs.rootDirectory, dirPath)
	if parentDir == nil {
		return fmt.Errorf("parent directory does not exist")
	}
	if _, exists := parentDir.ChildFiles[fileName]; exists {
		return fmt.Errorf("file already exists")
	}

	newFileInode := &utils.Inode{
		ID:        utils.GenerateInodeID(),
		Name:      fileName,
		IsDir:     false,
		Size:      0, // Size could be set appropriately
		Timestamp: time.Now(),
	}
	parentDir.ChildFiles[fileName] = newFileInode
	persistence.RecordEditLog("CREATE_FILE", filePath, newFileInode)

	return nil
}

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
func (fs *FileSystemService) CreateDirectory(dirPath string) error {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()
	parentPath, dirName := filepath.Dir(dirPath), filepath.Base(dirPath)
	parentDir := utils.FindDirectory(fs.rootDirectory, parentPath)
	if parentDir == nil {
		return fmt.Errorf("No parent path provided " + parentPath)
	}
	if _, exists := parentDir.ChildDirs[dirName]; exists {
		return fmt.Errorf("directory already exists")
	}

	newDirInode := &utils.Inode{
		ID:        utils.GenerateInodeID(),
		Name:      dirName,
		IsDir:     true,
		Timestamp: time.Now(),
	}
	parentDir.ChildDirs[dirName] = &utils.Directory{
		Inode:      newDirInode,
		ChildFiles: make(map[string]*utils.Inode),
		ChildDirs:  make(map[string]*utils.Directory),
	}

	persistence.RecordEditLog("CREATE_DIRECTORY", dirPath, newDirInode)

	return nil
}

func (fs *FileSystemService) ReadDirectory(dirPath string) ([]*utils.Inode, error) {
	fs.rootMutex.Lock()
	defer fs.rootMutex.Unlock()

	parentPath, dirName := filepath.Dir(dirPath), filepath.Base(dirPath)

	// Find the parent directory
	parentDir := utils.FindDirectory(fs.rootDirectory, parentPath)
	if parentDir == nil {
		return nil, fmt.Errorf("no parent path provided: %s", parentPath)
	}

	// Check if the directory exists
	dir, exists := parentDir.ChildDirs[dirName]
	if !exists {
		return nil, fmt.Errorf("directory does not exist: %s", dirPath)
	}

	// Read child files and directories
	childFiles := make([]*utils.Inode, 0, len(dir.ChildFiles))
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

func (fs *FileSystemService) AllocateFileBlocks(filePath string, fileSize int64) ([]utils.BlockAssignment, error) {
	const blockSize int64 = 64 * 1024 * 1024
	var blockAssignments []utils.BlockAssignment

	// dataNodeManager := gRPC.GetInstance()
	// dataNodes := dataNodeManager.GetDataNodes()
	// Simulate DataNode addresses

	// Calculate the number of blocks needed
	numBlocks := fileSize / blockSize
	if fileSize%blockSize != 0 {
		numBlocks++
	}

	// Assign blocks to DataNodes
	for i := int64(0); i < numBlocks; i++ {
		blockID := fmt.Sprintf("%s-block-%d", filepath.Base(filePath), i)
		// dataNodeIndex := i % int64(len(dataNodes)) // Simple round-robin allocation

		blockAssignments = append(blockAssignments, utils.BlockAssignment{
			BlockID: blockID,
			// DataNodeAddresses: []string{dataNodes[dataNodeIndex].Address},
		})
	}

	// return utils.AllocateFileBlocksResponse{BlockAssignments: blockAssignments}, nil
	return nil, nil
}
