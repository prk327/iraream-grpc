package handler

import (
	"context"
	"fmt"
	pb "gen-db-api/api/gen"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyService) GetData(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// Business logic, for now returning a static response
	return &pb.Response{
		Data: fmt.Sprintf("Hello, %s!", req.Query),
	}, nil
}
