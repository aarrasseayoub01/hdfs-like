package grpc

import (
	"context"

	"github.com/aarrasseayoub01/namenode/protobuf"
)

// NameNodeServer implements the protobuf-defined gRPC server interface
type DataNodeServer struct {
	protobuf.UnimplementedNameNodeServiceServer
}

// NewNameNodeServer creates a new instance of NameNodeServer
func NewDataNodeServer() *DataNodeServer {
	return &DataNodeServer{}
}

// RegisterDataNode is an example method implementing a service method
func (s *DataNodeServer) StoreBlock(ctx context.Context, in *protobuf.StoreBlockRequest) (*protobuf.StoreBlockResponse, error) {
	// Implement the logic here
	return &protobuf.StoreBlockResponse{Success: true}, nil
}
