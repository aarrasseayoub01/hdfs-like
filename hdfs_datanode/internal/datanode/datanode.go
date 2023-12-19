package datanode

import (
	"log"

	"github.com/aarrasseayoub01/namenode/datanode/internal/config"
)

type DataNode struct {
	config *config.Config
	// Add other fields as needed
}

func NewDataNode(cfg *config.Config) (*DataNode, error) {
	dn := &DataNode{
		config: cfg,
	}
	// Additional initialization here
	return dn, nil
}

func (dn *DataNode) Start() error {
	// Start the DataNode functionality
	log.Printf("Starting DataNode on %s", dn.config.DataNodeAddress)
	// Add more startup logic here
	return nil
}
