package main

import (
	"RPCing/echo"
	"RPCing/server"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Start the gRPC echo server in a separate goroutine
	// ( This lets the server run concurrently with the client in the same process )
	go server.StartEchoServer()

	// Create a root context for RPCs
	ctx := context.Background()

	// Establish a gRPC client connection to the server at localhost:8080
	// ( Using insecure credentials (no TLS) for local testing )
	conn, err := grpc.NewClient("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	// If the connection fails, panic
	if err != nil {
		panic(err)
	}

	// Ensure the connection is closed
	defer conn.Close()

	// Create a new Echo service client
	e := echo.NewEchoServerClient(conn)

	// Echo RPC call -> sending a message to the server
	resp, err := e.Echo(ctx, &echo.EchoRequest{
		Message: "Hello world, Testing echo server",
	})

	// If the RPC returns an error, panic
	if err != nil {
		panic(err)
	}

	fmt.Println("Server: ", resp.Response)
}
