package main

import (
	"context"
	"log"
	"time"

	hdfs "github.com/aarrasseayoub01/namenode/protobuf/hdfs"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("namenode:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hdfs.NewNameNodeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RegisterDataNode(ctx, &hdfs.RegisterDataNodeRequest{DatanodeAddress: "datanode_address"})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Registration Success: %t", r.GetSuccess())
}
