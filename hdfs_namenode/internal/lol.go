package main

import (
	"context"
	"log"
	"net"

	hdfs "github.com/aarrasseayoub01/namenode/protobuf/hdfs"
	"google.golang.org/grpc"
)

type nameNodeServer struct {
	hdfs.UnimplementedNameNodeServiceServer
}

func (s *nameNodeServer) RegisterDataNode(ctx context.Context, in *hdfs.RegisterDataNodeRequest) (*hdfs.RegisterDataNodeResponse, error) {
	// Implement your logic here
	return &hdfs.RegisterDataNodeResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	hdfs.RegisterNameNodeServiceServer(grpcServer, &nameNodeServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
