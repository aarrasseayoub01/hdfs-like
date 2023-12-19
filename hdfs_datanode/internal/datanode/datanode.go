package datanode

import (
	"fmt"
	"log"

	"github.com/aarrasseayoub01/namenode/datanode/internal/config"
	datamgmt "github.com/aarrasseayoub01/namenode/datanode/internal/datamngnt"
	"github.com/aarrasseayoub01/namenode/datanode/internal/grpc"
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
	if err := dn.startGRPCclient(); err != nil {
		return fmt.Errorf("failed to start gRPC client: %v", err)
	}

	log.Printf("Starting DataNode on %s", dn.config.DataNodeAddress)
	// Add more startup logic here
	return nil
}

func (dn *DataNode) startGRPCclient() error {
	nameNodeAddress := "localhost:50051" // Adjust the address of the NameNode
	client, err := grpc.NewDataNodeClient(nameNodeAddress)
	if err != nil {
		log.Fatalf("Failed to create DataNode client: %v", err)
		return err
	}
	defer client.Close()

	// Example: Register with NameNode
	if err := client.RegisterWithNameNode(); err != nil {
		log.Fatalf("Failed to register with NameNode: %v", err)
		return err
	}
	return nil
}
