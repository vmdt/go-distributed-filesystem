package cluster

import "google.golang.org/grpc"

type Node interface {
	CreateConn() (*grpc.ClientConn, error)
	IsConnAlive() bool
}

// Transport is anything that handles the communication between node with grpc
type Transport interface{}
