# Use Golang image to build the Go binary
FROM golang:1.19-alpine as builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Install dependencies and build the Go app
RUN go mod tidy
RUN go build -o grpc-api cmd/server/main.go

# Second stage: Build minimal Docker image with Go binary
FROM alpine:latest

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/grpc-api .

# Expose port
EXPOSE 50051

# Command to run the server
CMD ["./grpc-api"]
