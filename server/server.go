package server

import (
	"fmt"
	"net"
	"os"
)

func Start(address string) error {
	fmt.Println("Starting notRedis Server!...")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to bind port 6379")
		os.Exit(1)
	}

	fmt.Println("Server listening on Port 6379")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(connection)

	}

}
