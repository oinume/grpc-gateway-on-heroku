package main

import (
	"fmt"
	"log"
	"net"
	"os"

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
	defaultGRPCPort    = "50051"
	defaultGatewayPort = "50052"
)

func startGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Println("Starting gRPC server on " + port)
	return s.Serve(lis)
}

func main() {
	grpcPort := os.Getenv("GPRC_PORT")
	gatewayPort := os.Getenv("PORT")
	if grpcPort == "" {
		grpcPort = defaultGRPCPort
	}
	if gatewayPort == "" {
		gatewayPort = defaultGatewayPort
	}
	if grpcPort == gatewayPort {
		log.Fatalf("Can't specify same port.")
	}

	if err := startGRPCServer(grpcPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
