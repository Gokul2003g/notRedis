package server

import (
	"fmt"
	"net"
)

func handleConnection(connection net.Conn) {
	fmt.Println("Client connected: ", connection.RemoteAddr().String())

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error Reading input: ", err.Error())
			return
		}

		fmt.Println(string(buffer[:bytesRead]))
		connection.Write([]byte("+PONG\r\n"))
	}

}
