package main

import (
	"RPCing/chat"
	sc "RPCing/server/connection"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// ChatServer implements the gRPC Chat service and manages client connections
type ChatServer struct {
	broadcast   chan *chat.ChatMessage // Channel for broadcasting messages to all clients
	quit        chan struct{}          // Channel used to signal server shutdown
	connections []*sc.Connection       // Slice of active client connections
	lock        sync.Mutex             // Mutex for synchronizing access to connections
}

// NewChatServer constructs a ChatServer, starts its broadcaster loop, and returns it
func NewChatServer() *ChatServer {
	srv := &ChatServer{
		broadcast: make(chan *chat.ChatMessage),
		quit:      make(chan struct{}),
	}
	go srv.start()
	return srv
}

// Close signals the server to stop the broadcasting loop
func (c *ChatServer) Close() error {
	close(c.quit)
	return nil
}

// start listens for new messages or quit signal and handles broadcasting
func (c *ChatServer) start() {
	running := true
	for running {
		select {

		case msg := <-c.broadcast:
			// When message received, forward it to all connections
			c.lock.Lock()
			for _, v := range c.connections {
				go v.Send(msg)
			}
			c.lock.Unlock()

		case <-c.quit:
			// Stop the loop when quit signal is received
			running = false
		}
	}
}

// Chat is the gRPC method that handles a bidirectional stream connection from a client
func (c *ChatServer) Chat(stream chat.Chat_ChatServer) error {
	conn := sc.NewConnection(stream)

	// Add the new connection to the list
	c.lock.Lock()
	c.connections = append(c.connections, conn)
	c.lock.Unlock()

	// Start receiving messages
	// GetMessages is a blocking call that only returns when the client’s stream ends
	// ( because the client closed the stream, or an error occurred )
	err := conn.GetMessages(c.broadcast)

	// Remove the connection from the list
	// ( after client closed stream )
	c.lock.Lock()
	for i, v := range c.connections {
		if v == conn {
			c.connections = append(c.connections[:i], c.connections[i+1:]...)
			break
		}
	}
	c.lock.Unlock()

	return err
}

func main() {
	// Listen on TCP port 8080
	lstn, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Create gRPC server instance
	s := grpc.NewServer()
	// Create ChatServer instance and register it
	srv := NewChatServer()
	chat.RegisterChatServer(s, srv)

	// Start serving incoming connections
	fmt.Println("Server listening on port 8080")
	err = s.Serve(lstn)
	if err != nil {
		panic(err)
	}
}
