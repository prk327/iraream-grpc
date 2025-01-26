package service

import (
	pb "gen-db-api/api/gen"
	"gen-db-api/pkg/handler"

	"google.golang.org/grpc"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func RegisterService(grpcServer *grpc.Server) {
	pb.RegisterMyServiceServer(grpcServer, &handler.MyService{})
}
