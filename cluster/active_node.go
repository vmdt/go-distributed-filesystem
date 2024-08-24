package cluster

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var IPAddressTable = []string{
	"192.168.0.4:3000",
	"192.168.0.5:3000",
	"192.168.0.6:3000",
}

// Active Node represents the node in cluster
type ActiveNode struct {
	ListenAddress string
	Conn          *grpc.ClientConn
}

// NodesChecker is anything that keeps track of active nodes
type NodesChecker struct {
	Nodes []*ActiveNode
}

func NewActiveNode(listenAdrr string) *ActiveNode {
	return &ActiveNode{
		ListenAddress: listenAdrr,
	}
}

func (node *ActiveNode) IsConnAlive() bool {
	return true
}

func (node *ActiveNode) CreateConn() error {
	conn, err := grpc.NewClient(node.ListenAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	fmt.Printf("new grpc client [%v]\n", conn.Target())
	node.Conn = conn
	return nil
}

func (checker *NodesChecker) ReadAvailableNodes() error {
	for _, ip := range IPAddressTable {
		node := NewActiveNode(ip)
		err := node.CreateConn()

		if err != nil {
			log.Printf("create grpc connection error: %v\n", err)
			return err
		}

		checker.Nodes = append(checker.Nodes, node)
	}
	return nil
}
