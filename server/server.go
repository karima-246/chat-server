package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	// Read in the network port to listen on, from the commandline argument (default port 8030)
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, _ := net.Listen("tcp", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	clientId := 0

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client

			for {
				clients[clientId] = conn
				go handleClient(conn, clientId, msgs)
				clientId++
			}

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender

			reader := bufio.NewReader(os.Stdin)
			for i, _ := range clients {
				if i != msg.sender {
					msg, _ := reader.ReadString('\n')
					fmt.Fprint(msgs, msg) // ?
				}
			}
		}
	}
}
