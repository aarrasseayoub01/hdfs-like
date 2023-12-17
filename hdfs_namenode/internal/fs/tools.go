package fs

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FindDirectory(root *Directory, path string) *Directory {
	// Handle special case where the path is the root directory
	if path == "/" || path == "\\" || path == "" {
		return root
	}

	// Clean up the path and split it into components
	cleanedPath := filepath.Clean(path)[1:]
	parts := strings.Split(cleanedPath, string(os.PathSeparator))

	// Start from the root directory
	currentDir := root

	// Traverse the directory tree
	for _, part := range parts {
		if nextDir, ok := currentDir.ChildDirs[part]; ok {
			// Move to the next directory in the path
			currentDir = nextDir
		} else {
			// Directory not found in the path
			return nil
		}
	}

	return currentDir
}

func GenerateInodeID() int64 {
	return time.Now().UnixNano()
}
