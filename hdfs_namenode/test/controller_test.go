package service

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/aarrasseayoub01/hdfs-mini/internal/controller"
)

// Mock the FileSystemService interface
type MockFileSystemService struct {
	mock.Mock
}

func (m *MockFileSystemService) CreateFile(filePath string) error {
	args := m.Called(filePath)
	return args.Error(0)
}

func (m *MockFileSystemService) DeleteFile(filePath string) error {
	args := m.Called(filePath)
	return args.Error(0)
}

func (m *MockFileSystemService) CreateDirectory(dirPath string) error {
	args := m.Called(dirPath)
	return args.Error(0)
}

func (m *MockFileSystemService) DeleteDirectory(dirPath string) error {
	args := m.Called(dirPath)
	return args.Error(0)
}

func TestFileSystemController_CreateFileHandler(t *testing.T) {
	// Create a mock FileSystemService
	mockService := new(MockFileSystemService)
	controller := &controller.FileSystemController{Service: mockService}

	// Prepare a test request
	requestBody := `{"filePath": "/test.txt"}`
	req := httptest.NewRequest("POST", "/createFile", strings.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Mock the service method and call the handler
	mockService.On("CreateFile", "/test.txt").Return(nil)
	controller.CreateFileHandler(w, req)

	// Check the response status code and service method calls
	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestFileSystemController_DeleteFileHandler(t *testing.T) {
	// Create a mock FileSystemService
	mockService := new(MockFileSystemService)
	controller := &controller.FileSystemController{Service: mockService}

	// Prepare a test request
	requestBody := `{"filePath": "/test.txt"}`
	req := httptest.NewRequest("DELETE", "/deleteFile", strings.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Mock the service method and call the handler
	mockService.On("DeleteFile", "/test.txt").Return(nil)
	controller.DeleteFileHandler(w, req)

	// Check the response status code and service method calls
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestFileSystemController_CreateDirectoryHandler(t *testing.T) {
	// Create a mock FileSystemService
	mockService := new(MockFileSystemService)
	controller := &controller.FileSystemController{Service: mockService}

	// Prepare a test request
	requestBody := `{"dirPath": "/testdir"}`
	req := httptest.NewRequest("POST", "/createDirectory", strings.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Mock the service method and call the handler
	mockService.On("CreateDirectory", "/testdir").Return(nil)
	controller.CreateDirectoryHandler(w, req)

	// Check the response status code and service method calls
	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestFileSystemController_DeleteDirectoryHandler(t *testing.T) {
	// Create a mock FileSystemService
	mockService := new(MockFileSystemService)
	controller := &controller.FileSystemController{Service: mockService}

	// Prepare a test request
	requestBody := `{"dirPath": "/testdir"}`
	req := httptest.NewRequest("DELETE", "/deleteDirectory", strings.NewReader(requestBody))
	w := httptest.NewRecorder()

	// Mock the service method and call the handler
	mockService.On("DeleteDirectory", "/testdir").Return(nil)
	controller.DeleteDirectoryHandler(w, req)

	// Check the response status code and service method calls
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
