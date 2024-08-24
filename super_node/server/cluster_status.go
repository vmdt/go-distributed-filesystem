package server

import (
	"context"
	"time"

	fileservice "github.com/vmdt/distributed-filestorage/super_node/proto"
	"google.golang.org/grpc"
)

const (
	MinimumAvg float32 = 301.00
)

type Node struct {
	ListenAddress string
	ClusterName   string
}

type ClusterStatus struct {
	clusters map[string]string
}

func NewClusterStatus(clusterList map[string]string) *ClusterStatus {
	return &ClusterStatus{
		clusters: clusterList,
	}
}

// This method is responsible for finding 2 least loaded leader nodes from 2 clusters.
// First cluster is responsible for the primary cluster, another is replication
func (c *ClusterStatus) LeastUtilizedNode() (*Node, *Node) {
	minVal, minVal2 := MinimumAvg, MinimumAvg
	node1, node2 := &Node{ListenAddress: "", ClusterName: ""}, &Node{ListenAddress: "", ClusterName: ""}

	for cluster, node := range c.clusters {
		conn, err := c.isChannelAlive(node)

		if err == nil {
			defer conn.Close()
			fileClient := fileservice.NewFileServiceClient(conn)
			stats, _ := fileClient.GetClusterStats(context.Background(), &fileservice.Empty{})

			total := 300.00 - (stats.GetCpuUsage() + stats.GetDiskSpace() + stats.GetUsedMem())
			avgUsage := total / 3

			if avgUsage < minVal {
				minVal2 = minVal
				minVal = avgUsage
				node1.ClusterName = cluster
				node1.ListenAddress = node
			} else if avgUsage < minVal2 {
				minVal2 = avgUsage
				node2.ClusterName = cluster
				node2.ListenAddress = node
			}
		}
	}

	if node1.ListenAddress == "" && node2.ClusterName == "" {
		return nil, nil
	}

	return node1, node2
}

func (c *ClusterStatus) isChannelAlive(ipAdress string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, ipAdress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
