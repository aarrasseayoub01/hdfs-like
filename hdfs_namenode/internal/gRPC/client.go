package grpc

import (
	"context"
	"log"
	"time"

	"github.com/aarrasseayoub01/namenode/protobuf"
	"google.golang.org/grpc"
)

// NameNodeClient is a client for interacting with the DataNode gRPC service
type NameNodeClient struct {
	conn   *grpc.ClientConn
	client protobuf.DataNodeServiceClient
}

// NewNameNodeClient creates a new client for the DataNode service
func NewNameNodeClient(dataNodeAddress string) (*NameNodeClient, error) {
	conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := protobuf.NewDataNodeServiceClient(conn)
	return &NameNodeClient{conn: conn, client: client}, nil
}

// StoreBlock sends a StoreBlock request to a DataNode
func (c *NameNodeClient) StoreBlock(blockID string, blockData []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example: Sending a StoreBlockRequest to the DataNode
	response, err := c.client.StoreBlock(ctx, &protobuf.StoreBlockRequest{
		BlockId:   blockID,
		BlockData: blockData,
	})
	if err != nil {
		return err
	}
	log.Printf("StoreBlock response from DataNode: %v", response.GetSuccess())
	return nil
}

// Close closes the client connection
func (c *NameNodeClient) Close() {
	c.conn.Close()
}
