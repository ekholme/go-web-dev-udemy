package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
		go handle(conn)
	}

}

//handler function to use in our goroutine above
func handle(conn net.Conn) {

	//set a deadline on the connection
	//this will close the connection after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Conn Timeout!")
	}

	//create a new scanner using the connection
	scanner := bufio.NewScanner(conn)

	//for all of the lines in whatever we're reading, get the text (which will break with line breaks, i think), and return it.
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}

	fmt.Println("Now we done")

	defer conn.Close()
}

//to see this work, go to the terminal and connect to telnet localhost 8080
//then start typing in some messages
