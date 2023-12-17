package main

import (
	"fmt"
	"hdfs-datanode/internal/config"
	"hdfs-datanode/internal/datanode"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Error loading configuration: %v", err))
	}

	// Initialize DataNode
	dn, err := datanode.NewDataNode(cfg)
	if err != nil {
		panic(fmt.Sprintf("Error initializing DataNode: %v", err))
	}

	// Start DataNode
	err = dn.Start()
	if err != nil {
		panic(fmt.Sprintf("Error starting DataNode: %v", err))
	}
}
