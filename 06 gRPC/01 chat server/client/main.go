package main

import (
	"RPCing/chat"
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Ensure the user provides the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Error: expected 2 arguments.\nUsage: program <url> <username>")
		return
	}

	// Create a background context for the gRPC call
	ctx := context.Background()

	// Establish a gRPC client connection to the server at localhost:8080
	// ( Using insecure credentials (no TLS) for local testing )
	conn, err := grpc.NewClient("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a new Chat client
	c := chat.NewChatClient(conn)

	// Open a bidirectional stream for sending and receiving messages
	stream, err := c.Chat(ctx)
	if err != nil {
		panic(err)
	}

	// A channel of empty structs, often used when there's no need to send any actual data — just a signal
	// It's unbuffered, so it will block until someone sends (or closes) the channel
	waits := make(chan struct{})

	// Goroutine for receiving and printing messages from the server
	go func() {
		for {
			// Receive message from the server
			msg, err := stream.Recv()

			// If server has closed the stream
			if err == io.EOF {
				// Signal that receiving is done
				close(waits)
				return
			} else if err != nil {
				panic(err)
			}
			fmt.Println(msg.User, ":", msg.Message)
		}
	}()
	fmt.Println("Connection established\nType 'quit' to close the connection")

	// Read user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		// If user types "quit" -> close and exit
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}
		// Send the message to the server
		err := stream.Send(&chat.ChatMessage{
			User:    os.Args[2],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}
	// Block here until the receiving goroutine signals it's done (by closing the channel)
	<-waits
}
