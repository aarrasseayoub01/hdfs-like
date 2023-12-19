package persistence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/aarrasseayoub01/namenode/namenode/internal/fs"
)

type EditLogEntry struct {
	Timestamp time.Time
	Action    string
	Inode     *fs.Inode
	Path      string
}

const editLogFileName = "editlog.json"

var (
	editLog      []EditLogEntry
	editLogMutex sync.Mutex
)

func RecordEditLog(action string, path string, inode *fs.Inode) {
	// editLogMutex.Lock()
	// defer editLogMutex.Unlock()
	entry := EditLogEntry{
		Timestamp: time.Now(),
		Action:    action,
		Path:      path,
		Inode:     inode,
	}
	editLog = append(editLog, entry)
	if shouldTriggerCheckpoint() {
		triggerCheckpoint()
	}
	saveEditLog()
}

// Usage: recordEditLog("CREATE", inode)

func replayEditLog(root *fs.Directory) {
	for _, entry := range editLog {
		// Split the path to get parent directory and target name
		dirPath, targetName := filepath.Split(entry.Path)

		// Find the target directory based on dirPath
		targetDir := fs.FindDirectory(root, dirPath)

		switch entry.Action {
		case "CREATE_FILE":
			targetDir.ChildFiles[targetName] = entry.Inode

		case "DELETE_FILE":
			delete(targetDir.ChildFiles, targetName)

		case "CREATE_DIRECTORY":
			targetDir.ChildDirs[targetName] = &fs.Directory{
				Inode:      entry.Inode,
				ChildFiles: make(map[string]*fs.Inode),
				ChildDirs:  make(map[string]*fs.Directory),
			}

		case "DELETE_DIRECTORY":
			delete(targetDir.ChildDirs, targetName)

		}
	}
}

func loadEditLog() {
	// Check if the editlog file exists
	if _, err := os.Stat(editLogFileName); os.IsNotExist(err) {
		// Editlog file does not exist, create an empty editlog
		editLog = []EditLogEntry{}
		return
	}

	// Read the editlog file
	data, err := ioutil.ReadFile(editLogFileName)
	if err != nil {
		fmt.Printf("Error reading editlog file: %v\n", err)
		return
	}

	// Deserialize the editlog
	err = json.Unmarshal(data, &editLog)
	if err != nil {
		fmt.Printf("Error decoding editlog data: %v\n", err)
		return
	}
}

func saveEditLog() {
	// Serialize the editlog
	data, err := json.MarshalIndent(editLog, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding editlog data: %v\n", err)
		return
	}

	// Write the editlog to the file
	err = ioutil.WriteFile(editLogFileName, data, 0644)
	if err != nil {
		fmt.Printf("Error writing editlog file: %v\n", err)
		return
	}

	// If the editlog is empty, delete the editlog file
	if len(editLog) == 0 {
		err := os.Remove(editLogFileName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Printf("Error deleting editlog file: %v\n", err)
		}
	}
}
