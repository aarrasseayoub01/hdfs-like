package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	utils "github.com/aarrasseayoub01/hdfs-mini/internal/fs"
	svc "github.com/aarrasseayoub01/hdfs-mini/internal/service"
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

	err := c.Service.CreateFile(request.FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *FileSystemController) ReadFileHandler(w http.ResponseWriter, r *http.Request) {
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

	err := c.Service.CreateDirectory(request.DirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
