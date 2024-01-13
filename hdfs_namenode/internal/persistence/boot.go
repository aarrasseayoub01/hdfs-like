package persistence

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aarrasseayoub01/namenode/namenode/internal/fs"
)

var rootDirectory *fs.Directory

func InitializeFileSystem() *fs.Directory {
	fsImageExists := checkFsImageExists("fsimage.gob")

	if fsImageExists {
		var err error
		rootDirectory, err = loadFsImage("fsimage.gob")
		if err != nil {
			log.Fatalf("Failed to load filesystem image: %v", err)
		}
		loadEditLog()
		replayEditLog(rootDirectory)
	} else {
		rootDirectory = &fs.Directory{
			Inode: &fs.Inode{
				ID:        1, // root directory ID, usually 1
				Name:      "/",
				IsDir:     true,
				Size:      0,
				Blocks:    nil,
				Timestamp: time.Now(),
			},
			ChildFiles: make(map[string]*fs.Inode),
			ChildDirs:  make(map[string]*fs.Directory),
		}
	}
	return rootDirectory
}

func checkpointFileSystem() {

	err := saveFsImage(rootDirectory, "fsimage.gob")
	if err != nil {
		log.Printf("Error during filesystem checkpoint: %v", err)
	}
}

func triggerCheckpoint() error {

	editLogMutex.Lock()
	defer editLogMutex.Unlock()

	// Save the current state of the filesystem
	err := saveFsImage(rootDirectory, "fsimage.gob")
	if err != nil {
		return fmt.Errorf("failed to save FsImage: %w", err)
	}

	// Clear the edit log
	editLog = []EditLogEntry{}

	return nil
}

func checkFsImageExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	return err == nil // File exists and no error occurred
}
