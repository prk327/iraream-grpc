package main

import (
	pb "gen-db-api/api/gen" // Import generated gRPC code
	"gen-db-api/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Listen on TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the service
	pb.RegisterMyServiceServer(grpcServer, &service.MyService{})

	// Start serving
	log.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
