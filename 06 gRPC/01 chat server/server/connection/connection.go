package connection

import (
	"RPCing/chat"
	"io"
)

// Connection wraps a gRPC Chat_ChatServer stream and manages send/receive logic.
// It maintains an internal send queue and a quit signal channel.
// Messages sent by the client are forwarded to the broadcast channel.
type Connection struct {
	conn chat.Chat_ChatServer   // Underlying gRPC stream
	send chan *chat.ChatMessage // Channel to queue outbound messages
	quit chan struct{}          // Channel to signal shutdown
}

// NewConnection initializes a new Connection.
// It starts the send loop in a separate goroutine.
func NewConnection(conn chat.Chat_ChatServer) *Connection {
	c := &Connection{
		conn: conn,
		send: make(chan *chat.ChatMessage),
		quit: make(chan struct{}),
	}
	go c.start()
	return c
}

// Close signals the connection to stop and closes channels.
func (c *Connection) Close() error {
	close(c.quit)
	close(c.send)
	return nil
}

// Send enqueues a ChatMessage for sending to the client stream.
// If the connection is closed, the send attempt is ignored via recover().
func (c *Connection) Send(msg *chat.ChatMessage) {
	defer func() {
		recover()
	}()
	c.send <- msg
}

// start forwards messages from the channel to the gRPC stream until 'quit' signal is received.
// (runs in a goroutine)
func (c *Connection) start() {
	running := true
	for running {
		select {

		case msg := <-c.send:
			// Send the queued message over the gRPC stream
			c.conn.Send(msg)

		case <-c.quit:
			// Received shutdown signal
			running = false
		}
	}
}

// GetMessages reads incoming messages from the client stream.
// Each received message is forwarded to the broadcast channel in its own goroutine.
// On EOF, it closes the connection and returns the error.
func (c *Connection) GetMessages(broadcast chan<- *chat.ChatMessage) error {
	for {
		// Block until a message arrives or an error occurs
		msg, err := c.conn.Recv()

		if err == io.EOF {
			c.Close()
			return err
		}

		// Forward the message to the broadcast channel
		// (in a Goroutine, without blocking for loop)
		go func(msg *chat.ChatMessage) {
			select {
			case broadcast <- msg:
			case <-c.quit:
			}
		}(msg)
	}
}
