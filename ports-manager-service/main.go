package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"ports-manager-service/api"
	pb "ports-manager-service/grpc"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPortManagerServer(grpcServer, api.NewServer())

	log.Printf("running server on port %d", *port)
	log.Fatal(grpcServer.Serve(lis))
}
