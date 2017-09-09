package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oinume/grpc-gateway-on-heroku/gen/go/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *echo.StringMessage) (*echo.StringMessage, error) {
	return &echo.StringMessage{Value: "Echo:" + in.Value}, nil
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Println("Starting a server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
