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
	mu        sync.Mutex
	dataNodes map[string]*DataNode
}

// NewNameNodeServer creates a new instance of NameNodeServer
func NewNameNodeServer() *NameNodeServer {
	return &NameNodeServer{
		dataNodes: make(map[string]*DataNode),
	}
}

func (s *NameNodeServer) RegisterDataNode(ctx context.Context, req *protobuf.RegisterDataNodeRequest) (*protobuf.RegisterDataNodeResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	address := req.GetDatanodeAddress()
	log.Printf("Registering DataNode with address: %s", address)

	id := uuid.New().String()
	var datanodeID string
	if existingNode, exists := s.dataNodes[address]; !exists {
		datanodeID = id // Generate a unique ID
		s.dataNodes[address] = &DataNode{
			Address: address,
			ID:      datanodeID,
		}
		// Additional initialization for the DataNode struct
	} else {
		datanodeID = existingNode.ID
		// Handle the case where the DataNode is already registered
	}

	return &protobuf.RegisterDataNodeResponse{Success: true, DatanodeId: datanodeID}, nil
}

func (s *NameNodeServer) SendHeartbeat(ctx context.Context, req *protobuf.HeartbeatRequest) (*protobuf.HeartbeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	datanodeID := req.GetDatanodeId()
	log.Printf("Received heartbeat from DataNode: %s", datanodeID)

	if _, exists := s.dataNodes[datanodeID]; exists {
		// Update the existing DataNode record with new heartbeat information
	} else {
		// Optionally handle the case where a heartbeat is received from an unknown DataNode
	}

	return &protobuf.HeartbeatResponse{Success: true}, nil
}
