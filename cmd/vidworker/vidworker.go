package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/MovAh13h/Vidsy/proto"
	"google.golang.org/grpc"
)

func main() {
	// take port from command line flags
	port := flag.Int("port", 3000, "Port number for the server")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterVidsyServer(grpcServer, &Server{})
	log.Printf("[Worker] Serving on port: %d", *port)
	grpcServer.Serve(lis)
}