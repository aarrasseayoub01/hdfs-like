package datanode

import (
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/aarrasseayoub01/namenode/datanode/internal/config"
	ctrl "github.com/aarrasseayoub01/namenode/datanode/internal/controller"
	datamgmt "github.com/aarrasseayoub01/namenode/datanode/internal/datamngnt"
	mng "github.com/aarrasseayoub01/namenode/datanode/internal/datamngnt" // Replace with your actual project path
	gRPC "github.com/aarrasseayoub01/namenode/datanode/internal/gRPC"
	"github.com/aarrasseayoub01/namenode/protobuf"
	"github.com/gorilla/mux"
)

type DataNode struct {
	config      *config.Config
	dataManager *datamgmt.DataManager

	// Add other fields as needed
}

func NewDataNode(cfg *config.Config) (*DataNode, error) {
	dm := datamgmt.NewDataManager("./") // Adjust the path as needed

	dn := &DataNode{
		config:      cfg,
		dataManager: dm,
	}
	// Additional initialization here
	return dn, nil
}

func (dn *DataNode) Start() error {

	r := mux.NewRouter()

	// Initialize DataManager
	dataManager := mng.NewDataManager("./") // Set the base directory path

	// Create a new Controller instance
	controller := ctrl.NewController(dataManager)
	// Define the routes
	r.HandleFunc("/addBlock", controller.AddBlock).Methods("POST")
	r.HandleFunc("/getBlock/{blockId}", controller.GetBlock).Methods("GET") // New route

	// Start the server
	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
	// Start the DataNode functionality
	go dn.startGRPCclient()

	go startGRPCserver()

	log.Printf("Starting DataNode on %s", dn.config.DataNodeAddress)
	// Add more startup logic here
	return nil
}

func (dn *DataNode) startGRPCclient() error {
	nameNodeAddress := "localhost:50051" // Adjust the address of the NameNode
	client, err := gRPC.NewDataNodeClient(nameNodeAddress)
	if err != nil {
		log.Fatalf("Failed to create DataNode client: %v", err)
		return err
	}
	defer client.Close()

	dataNodeID, err := client.RegisterWithNameNode()
	if err != nil {
		log.Fatalf("Failed to register with NameNode: %v", err)
	}

	go func(address string) {
		ticker := time.NewTicker(30 * time.Second) // Adjust the interval as needed
		nameNodeAddress := "localhost:50051"       // Adjust the address of the NameNode
		client, err := gRPC.NewDataNodeClient(nameNodeAddress)
		if err != nil {
			log.Fatalf("Failed to create DataNode client: %v", err)
		}
		defer client.Close()

		for range ticker.C {
			err := client.SendHeartbeat(address) // Use a unique identifier for the DataNode
			if err != nil {
				log.Printf("Error sending heartbeat: %v", err)
				// Handle error, maybe with a retry mechanism
			}
		}
	}(dataNodeID)

	return nil
}

func startGRPCserver() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterNameNodeServiceServer(grpcServer, gRPC.NewDataNodeServer())
	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
