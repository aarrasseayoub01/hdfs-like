package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aarrasseayoub01/hdfs-mini/internal/controller"
	"github.com/aarrasseayoub01/hdfs-mini/internal/persistence"
)

func main() {
	// Initialize the file system service
	rootDir := persistence.InitializeFileSystem()

	// Set up the controller with the service
	controller := controller.NewFileSystemController(rootDir)

	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/createFile", controller.CreateFileHandler).Methods("POST")
	r.HandleFunc("/readFile", controller.ReadFileHandler).Methods("GET")
	r.HandleFunc("/deleteFile", controller.DeleteFileHandler).Methods("DELETE")
	r.HandleFunc("/createDir", controller.CreateDirectoryHandler).Methods("POST")
	r.HandleFunc("/readDir", controller.ReadDirectoryHandler).Methods("GET")
	r.HandleFunc("/deleteDir", controller.DeleteDirectoryHandler).Methods("DELETE")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
