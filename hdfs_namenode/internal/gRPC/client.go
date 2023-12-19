package grpc

import (
	"context"
	"log"
	"time"

	"github.com/aarrasseayoub01/namenode/protobuf"
	"google.golang.org/grpc"
)

// DataNodeClient is a client for interacting with the NameNode gRPC service
type DataNodeClient struct {
	conn   *grpc.ClientConn
	client protobuf.NameNodeServiceClient
}

// NewDataNodeClient creates a new client for the NameNode service
func NewDataNodeClient(nameNodeAddress string) (*DataNodeClient, error) {
	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := protobuf.NewNameNodeServiceClient(conn)
	return &DataNodeClient{conn: conn, client: client}, nil
}

// RegisterWithNameNode registers the DataNode with the NameNode
func (c *DataNodeClient) RegisterWithNameNode() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example: Sending a RegisterDataNodeRequest to the NameNode
	response, err := c.client.RegisterDataNode(ctx, &protobuf.RegisterDataNodeRequest{DatanodeAddress: "datanode_address"})
	if err != nil {
		return err
	}
	log.Printf("Registration response from NameNode: %v", response.GetSuccess())
	return nil
}

// Close closes the client connection
func (c *DataNodeClient) Close() {
	c.conn.Close()
}
