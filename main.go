package main

import (
	"github.com/gokul2003g/notRedis/server"
	"log"
)

func main() {

	address := "0.0.0.0:6379"
	err := server.Start(address)

	if err != nil {
		log.Fatalln("Failed to Start notRedis Server. Error: ", err)
	}

}
