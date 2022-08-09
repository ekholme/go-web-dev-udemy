package main

import (
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
	}

	//RESUME AT 5:30 IN 22. TCP SERVER READ FROM CONNECTION VID
}
