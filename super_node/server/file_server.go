package server

import (
	"context"
	"fmt"
	"sync"

	fileservice "github.com/vmdt/distributed-filestorage/super_node/proto"
)

type FileServer struct {
	fileservice.UnimplementedFileServiceServer
	ListenAddress  string
	ClusterLeaders map[string]string
	mu             sync.Mutex
}

func NewFileServer(listenAddr string) *FileServer {
	return &FileServer{
		ListenAddress:  listenAddr,
		ClusterLeaders: make(map[string]string),
	}
}

// GetLeaderInfo gets invoked when each cluster's leader informs the super-node
// about who the current cluster leader is
func (s *FileServer) GetLeaderInfo(ctx context.Context, request *fileservice.ClusterInfo) (*fileservice.Ack, error) {
	fmt.Println("GetLeaderInfo called")

	address := request.Ip + ":" + request.Port
	s.mu.Lock()
	s.ClusterLeaders[request.ClusterName] = address
	s.mu.Unlock()

	return &fileservice.Ack{
		Success: true,
		Message: "Leader Updated.",
	}, nil
}
