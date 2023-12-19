package datanode

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/aarrasseayoub01/namenode/datanode/internal/config"
	datamgmt "github.com/aarrasseayoub01/namenode/datanode/internal/datamngnt"
	gRPC "github.com/aarrasseayoub01/namenode/datanode/internal/gRPC"
	"github.com/aarrasseayoub01/namenode/protobuf"
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

	// Start the DataNode functionality
	go dn.startGRPCclient()

	startGRPCserver()

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

	go func(id string) {
		ticker := time.NewTicker(30 * time.Second) // Adjust the interval as needed
		for range ticker.C {
			err := client.SendHeartbeat(id) // Use a unique identifier for the DataNode
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
