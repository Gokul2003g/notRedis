package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting notRedis Server!...")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind port 6379")
		os.Exit(1)
	}

	fmt.Println("Server listening on port 6379")

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Client connected: ", connection.RemoteAddr().String())

	defer connection.Close()
}
