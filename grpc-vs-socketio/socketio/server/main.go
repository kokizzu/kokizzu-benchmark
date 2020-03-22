package main

import (
	"log"
	"net/http"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// Message -
type Message struct {
	Text string `json:"text"`
}

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		// log.Println("Connected")
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		// log.Println("Disconnected")
	})

	server.On("/message", func(c *gosocketio.Channel, msg Message) string {
		return "hello world"
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	// log.Println("Starting server on port 15001...")
	log.Panic(http.ListenAndServe(":15001", serveMux))
}
