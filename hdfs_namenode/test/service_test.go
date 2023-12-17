// file_system_service_test.go

package service

import (
	"testing"

	"github.com/aarrasseayoub01/hdfs-mini/internal/fs"
	"github.com/aarrasseayoub01/hdfs-mini/internal/persistence"
	"github.com/aarrasseayoub01/hdfs-mini/internal/service"
	"github.com/stretchr/testify/assert"
)

func setupMockFileSystem() *fs.Directory {
	return &fs.Directory{
		Inode: &fs.Inode{
			ID:    1,
			Name:  "/",
			IsDir: true,
		},
		ChildFiles: make(map[string]*fs.Inode),
		ChildDirs:  make(map[string]*fs.Directory),
	}
}

func TestCreateFile(t *testing.T) {
	rootDir := persistence.InitializeFileSystem()
	service := service.NewFileSystemService(rootDir)

	// Test creating a new file
	err := service.CreateFile("/testfile.txt")
	assert.NoError(t, err)

	// Test trying to create a file that already exists
	err = service.CreateFile("/testfile.txt")
	assert.Error(t, err)

	// Optionally, more assertions to verify the state of rootDir
}

func TestDeleteFile(t *testing.T) {
	rootDir := persistence.InitializeFileSystem()
	service := service.NewFileSystemService(rootDir)

	// Setup: create a file to delete
	_ = service.CreateFile("/testfile.txt")

	// Test deleting the file
	err := service.DeleteFile("/testfile.txt")
	assert.NoError(t, err)

	// Test trying to delete a non-existent file
	err = service.DeleteFile("/nonexistent.txt")
	assert.Error(t, err)
}

func TestCreateDirectory(t *testing.T) {
	rootDir := persistence.InitializeFileSystem()
	service := service.NewFileSystemService(rootDir)

	// Test creating a new directory
	err := service.CreateDirectory("/newdir")
	assert.NoError(t, err)

	// Test trying to create a directory that already exists
	err = service.CreateDirectory("/newdir")
	assert.Error(t, err)
}

func TestDeleteDirectory(t *testing.T) {
	rootDir := persistence.InitializeFileSystem()
	service := service.NewFileSystemService(rootDir)

	// Setup: create a directory to delete
	_ = service.CreateDirectory("/newdir")

	// Test deleting the directory
	err := service.DeleteDirectory("/newdir")
	assert.NoError(t, err)

	// Test trying to delete a non-existent directory
	err = service.DeleteDirectory("/nonexistentdir")
	assert.Error(t, err)
}
