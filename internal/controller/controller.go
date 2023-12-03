package controller

import (
	"encoding/json"
	"net/http"

	utils "github.com/aarrasseayoub01/hdfs-mini/internal/fs"
	svc "github.com/aarrasseayoub01/hdfs-mini/internal/service"
)

type FileSystemController struct {
	service *svc.FileSystemService
}

func NewFileSystemController(rootDir *utils.Directory) *FileSystemController {
	fileSystemService := svc.NewFileSystemService(rootDir)

	return &FileSystemController{service: fileSystemService}
}

func (c *FileSystemController) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		FilePath string `json:"filePath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.service.CreateFile(request.FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *FileSystemController) DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		FilePath string `json:"filePath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.service.DeleteFile(request.FilePath)
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

	err := c.service.CreateDirectory(request.DirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *FileSystemController) DeleteDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DirPath string `json:"dirPath"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.service.DeleteDirectory(request.DirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
