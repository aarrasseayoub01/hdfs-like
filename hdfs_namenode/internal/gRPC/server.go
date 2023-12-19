package grpc

import (
	"context"
	"log"
	"sync"

	"github.com/aarrasseayoub01/namenode/protobuf"
)

// NameNodeServer implements the protobuf-defined gRPC server interface
type NameNodeServer struct {
	protobuf.UnimplementedNameNodeServiceServer
	mu        sync.Mutex
	dataNodes map[string]struct{}
}

// NewNameNodeServer creates a new instance of NameNodeServer
func NewNameNodeServer() *NameNodeServer {
	return &NameNodeServer{
		dataNodes: make(map[string]struct{}),
	}
}

func (s *NameNodeServer) RegisterDataNode(ctx context.Context, req *protobuf.RegisterDataNodeRequest) (*protobuf.RegisterDataNodeResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	address := req.GetDatanodeAddress()
	log.Printf("Registering DataNode with address: %s", address)

	// Add the DataNode to the map
	s.dataNodes[address] = struct{}{}

	// You can add additional logic here, such as updating metadata, etc.

	return &protobuf.RegisterDataNodeResponse{Success: true}, nil
}
