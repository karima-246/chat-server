package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

// The client needs to
// (a) process user input and send it to the server on each newline;
// (b) at the same time, accept messages sent from the server and display them as they arrive.

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn) // read from connection
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin) // read from stdin (user input)
	for {
		fmt.Println("Enter text to send -> ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)

	//TODO Start asynchronously reading and displaying messages
	for {
		go read(conn)  // ?
		go write(conn) // ?
	}

	//TODO Start getting and sending user messages.

}
