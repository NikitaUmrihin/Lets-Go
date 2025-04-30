package server

import (
	"RPCing/echo"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Response: "{ECHO}" + req.Message}, nil
}

func StartEchoServer() {
	// Start listening on TCP port 8080
	lstn, err := net.Listen("tcp", ":8080")

	// Panic if the listener cannot be started (e.g., port already in use)
	if err != nil {
		panic(err)
	}

	// Create a new gRPC server instance
	s := grpc.NewServer()

	// Create an instance of EchoServer
	srv := &EchoServer{}

	// Register the EchoServer to the gRPC server
	// (This connects my implementation to the gRPC framework so it can handle requests)
	echo.RegisterEchoServerServer(s, srv)

	// Start serving incoming gRPC requests on the listener
	err = s.Serve(lstn)

	// Panic if the server fails to start
	if err != nil {
		panic(err)
	}

	fmt.Println("Echo server running on port 8080")
}
