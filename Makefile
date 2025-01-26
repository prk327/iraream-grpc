proto:
	protoc --go_out=api/gen --go-grpc_out=api/gen api/proto/service.proto

start:
	go run cmd/server/main.go

build:
	./scripts/docker-build.sh

run:
	./scripts/run.sh

test:
	go test ./...
