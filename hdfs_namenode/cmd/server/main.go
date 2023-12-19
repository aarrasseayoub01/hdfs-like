package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/aarrasseayoub01/namenode/namenode/internal/controller"
	grpc2 "github.com/aarrasseayoub01/namenode/namenode/internal/grpc"
	"github.com/aarrasseayoub01/namenode/namenode/internal/persistence"
	"github.com/aarrasseayoub01/namenode/protobuf"
)

func main() {
	// Start the REST server
	go startRESTserver()

	// Start the gRPC server
	startGRPCserver()
}

func startRESTserver() {
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

func startGRPCserver() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterNameNodeServiceServer(grpcServer, grpc2.NewNameNodeServer()) // Use the NewNameNodeServer function
	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
