package main

import (
	"fmt"
	"log"
	"net"

	fileservice "github.com/vmdt/distributed-filestorage/super_node/proto"
	"github.com/vmdt/distributed-filestorage/super_node/server"
	"google.golang.org/grpc"
)

func main() {
	host := "192.168.0.9"
	port := "9000"

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen on [%s]: %v", lis.Addr().String(), err)
	}

	grpcServer := grpc.NewServer()
	fileServer := server.NewFileServer(lis.Addr().String())
	fileservice.RegisterFileServiceServer(grpcServer, fileServer)

	log.Printf("supernode server is running at [%s]\n", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
