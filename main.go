package main

import (
	"log"

	"github.com/vmdt/distributed-filestorage/cluster"
)

func main() {
	cluster := &cluster.NodesChecker{}
	err := cluster.ReadAvailableNodes()

	if err != nil {
		log.Panicln(err)
	}

	for {
	}
}
