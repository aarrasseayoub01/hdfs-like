package gRPC

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/aarrasseayoub01/namenode/protobuf" // Adjust this import path to where your protobuf definitions are.

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

func (c *DataNodeClient) RegisterWithNameNode() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dataNodeAddress := getLocalIPAddress()

	response, err := c.client.RegisterDataNode(ctx, &protobuf.RegisterDataNodeRequest{DatanodeAddress: dataNodeAddress})
	if err != nil {
		return "", err
	}
	log.Printf("Registered with NameNode, assigned ID: %s", response.GetDatanodeId())
	return response.GetDatanodeId(), nil
}

func (c *DataNodeClient) SendHeartbeat(datanodeID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Sending a HeartbeatRequest to the NameNode
	response, err := c.client.SendHeartbeat(ctx, &protobuf.HeartbeatRequest{DatanodeId: datanodeID})
	if err != nil {
		return err
	}
	log.Printf("Heartbeat response from NameNode: %v", response.GetSuccess())
	return nil
}

// Close closes the client connection
func (c *DataNodeClient) Close() {
	c.conn.Close()
}

func getLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Failed to get IP addresses: %v", err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
