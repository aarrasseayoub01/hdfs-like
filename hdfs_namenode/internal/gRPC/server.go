package grpc

import (
	"context"

	"github.com/aarrasseayoub01/namenode/protobuf"
)

// NameNodeServer implements the protobuf-defined gRPC server interface
type NameNodeServer struct {
	protobuf.UnimplementedNameNodeServiceServer
}

// NewNameNodeServer creates a new instance of NameNodeServer
func NewNameNodeServer() *NameNodeServer {
	return &NameNodeServer{}
}

// RegisterDataNode is an example method implementing a service method
func (s *NameNodeServer) RegisterDataNode(ctx context.Context, in *protobuf.RegisterDataNodeRequest) (*protobuf.RegisterDataNodeResponse, error) {
	// Implement the logic here
	return &protobuf.RegisterDataNodeResponse{Success: true}, nil
}
