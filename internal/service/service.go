package service

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	utils "github.com/aarrasseayoub01/hdfs-mini/internal/fs"
	"github.com/aarrasseayoub01/hdfs-mini/internal/persistence"
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
		return fmt.Errorf(parentPath)
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
