package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/MovAh13h/Vidsy/proto"
	"google.golang.org/grpc"
)

func getGrpcClients(workers string) *[]proto.VidsyClient {
	// split the comma separated ips
	ips := strings.Split(workers, ",")

	// to store all gRPC clients
	var clients []proto.VidsyClient
	
	for _, ip := range ips {
		// TODO: Add Certs here
		conn, err := grpc.Dial(ip, grpc.WithInsecure())	
	
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		// TODO: is this correct place to defer?
		defer conn.Close()

		// store all gRPC clients
		clients = append(clients, proto.NewVidsyClient(conn))
	}

	return &clients
}

func main() {
	// port number
	port := flag.Int("port", 3000, "Port number for the server")
	
	// ips of workers
	workerFlag := flag.String("workers", "",
		"Comma separated address of workers")

	// parse command line flags
	flag.Parse()

	fmt.Println(*port)

	workers := getGrpcClients(*workerFlag)
}