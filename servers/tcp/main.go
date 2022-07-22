package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//create a listener
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	//close the listener
	defer li.Close()

	for {
		//accept the connection and return a connection
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		//we can do any of these things bc conn implements a writer interface
		//it also implements a reader interface
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "how are you doing?")
		fmt.Fprintf(conn, "%v", "Well I hope!")

		conn.Close()

		//note that when this is run, we can make requests from the shell via: telnet localhost 8080
	}
}
