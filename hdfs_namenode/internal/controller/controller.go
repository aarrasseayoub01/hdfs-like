package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	utils "github.com/aarrasseayoub01/namenode/namenode/internal/fs"
	svc "github.com/aarrasseayoub01/namenode/namenode/internal/service"
)

type FileSystemController struct {
	Service *svc.FileSystemService
}

func NewFileSystemController(rootDir *utils.Directory) *FileSystemController {
	fileSystemService := svc.NewFileSystemService(rootDir)

	return &FileSystemController{Service: fileSystemService}
}

func (c *FileSystemController) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		FilePath string `json:"filePath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileInode, err := c.Service.CreateFile(request.FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(fileInode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *FileSystemController) ReadFileHandler(w http.ResponseWriter, r *http.Request) {

	// file, err := os.Open("../../cmd/server/fsimage.gob")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// // Create a decoder
	// decoder := gob.NewDecoder(file)

	// // Decode the data
	// var data fs.Directory
	// err = decoder.Decode(&data)
	// if err != nil {
	// 	fmt.Println("Error decoding GOB data:", err)
	// 	return
	// }

	// // Use the data
	// fmt.Println("Decoded data:", data)
	// Parse the query from URL
	query := r.URL.RawQuery

	// Read file
	fileInode, err := c.Service.ReadFile(strings.Split(query, "=")[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the fileInode as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(fileInode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *FileSystemController) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		FilePath string `json:"filePath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.Service.DeleteFile(request.FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *FileSystemController) CreateDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DirPath string `json:"dirPath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inode, err := c.Service.CreateDirectory(request.DirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(inode); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}
func (c *FileSystemController) ReadDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query from URL
	query := r.URL.RawQuery

	// Read Directory
	inodes, err := c.Service.ReadDirectory(strings.Split(query, "=")[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the inodes as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(inodes)
	if err := json.NewEncoder(w).Encode(inodes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *FileSystemController) DeleteDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DirPath string `json:"dirPath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.Service.DeleteDirectory(request.DirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *FileSystemController) AllocateFileBlocksHandler(w http.ResponseWriter, r *http.Request) {
	var request utils.AllocateFileBlocksRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Call the service layer to get block assignments
	allocateFileBlocks, err := c.Service.AllocateFileBlocks(request.FilePath, request.FileSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error allocating file blocks: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the block assignments
	response := utils.AllocateFileBlocksResponse{BlockAssignments: allocateFileBlocks.BlockAssignments}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}
