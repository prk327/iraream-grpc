package grpc_test

import (
	"context"
	"gen-db-api/api/gen"
	"gen-db-api/pkg/service"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestGRPCServer(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()

	// Register the service
	gen.RegisterMyServiceServer(s, &service.MyService{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	// Create a client and make requests (mock client code)
	conn, err := grpc.Dial("bufnet", grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := gen.NewMyServiceClient(conn)

	// Call gRPC method
	resp, err := client.GetData(context.Background(), &gen.Request{Query: "test"})
	if err != nil {
		t.Fatalf("GetData failed: %v", err)
	}

	if resp.Data != "expected data" {
		t.Errorf("expected data, got %v", resp.Data)
	}
}
