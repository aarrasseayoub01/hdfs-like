package gRPC

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"

	"github.com/aarrasseayoub01/namenode/protobuf"
)

type DataNode struct {
	Address string
	ID      string
	// Add other relevant fields such as status, last heartbeat time, etc.
}

// NameNodeServer implements the protobuf-defined gRPC server interface
type NameNodeServer struct {
	protobuf.UnimplementedNameNodeServiceServer
}

// Package datanode_manager or a similar name

type DataNodeManager struct {
	mu        sync.RWMutex
	dataNodes map[string]*DataNode
}

var instance *DataNodeManager
var once sync.Once

// GetInstance returns the singleton instance of DataNodeManager
func GetInstance() *DataNodeManager {
	once.Do(func() {
		instance = &DataNodeManager{
			dataNodes: make(map[string]*DataNode),
		}
	})
	return instance
}

func NewNameNodeServer() *NameNodeServer {
	return &NameNodeServer{}
}
func (m *DataNodeManager) RegisterDataNode(address, id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dataNodes[address] = &DataNode{Address: address, ID: id}
}

func (m *DataNodeManager) GetDataNodes() map[string]*DataNode {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// Return a copy of the map to avoid concurrent modifications
	dataNodesCopy := make(map[string]*DataNode)
	for k, v := range m.dataNodes {
		dataNodesCopy[k] = v
	}
	return dataNodesCopy
}

// Other necessary methods...

func (s *NameNodeServer) RegisterDataNode(ctx context.Context, req *protobuf.RegisterDataNodeRequest) (*protobuf.RegisterDataNodeResponse, error) {
	address := req.GetDatanodeAddress()
	log.Printf("Registering DataNode with address: %s", address)

	datanodeID := uuid.New().String()
	dataNodeManager := GetInstance()
	dataNodeManager.RegisterDataNode(address, datanodeID)

	return &protobuf.RegisterDataNodeResponse{Success: true, DatanodeId: datanodeID}, nil
}

func (s *NameNodeServer) SendHeartbeat(ctx context.Context, req *protobuf.HeartbeatRequest) (*protobuf.HeartbeatResponse, error) {
	address := req.GetDatanodeAddress()

	log.Printf("Received heartbeat from DataNode: %s", address)

	dataNodeManager := GetInstance()
	dataNodes := dataNodeManager.GetDataNodes()

	// Check if the DataNode exists in the manager
	if dataNode, exists := dataNodes[address]; exists {
		// Update the existing DataNode record with new heartbeat information
		// This could include updating a timestamp, status, etc.
		// Note: Ensure thread-safe updating if necessary
		log.Printf("Heartbeat received from known DataNode: %s", dataNode.Address)
	} else {
		// Optionally handle the case where a heartbeat is received from an unknown DataNode
		log.Printf("Heartbeat received from unknown DataNode: %s", address)
	}

	return &protobuf.HeartbeatResponse{Success: true}, nil
}
