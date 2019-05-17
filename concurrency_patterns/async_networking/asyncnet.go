package main

import (
	"net"
	"log"
	"io"
	"os"
)

func serviceConnection(conn net.Conn) {
		// both os.Stderr and conn implements
		// Reader and Writer methods
		// io.Copy is blocking and stops the Accept loop
		// hence not serving other connections
		numOfBytes, err := io.Copy(os.Stderr, conn)
		log.Printf("Copied %d bytes; fineshed with err = %v", numOfBytes, err)
}

func main() {
	// net.Listen takes to arguments
	// on which protocol we want listen on, as string
	// address with a port number, as string
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// take the new connections that come in via tcp
	// and service them
	for {
		
		// make connection
		// process of making connection is blocking
		conn, err := listener.Accept()

		// check whether err implements the
		// net.Error interface
		// if the error is temporary, you can retry
		if err != nil {
			log.Fatal(err)
		}

		go serviceConnection(conn)

	}
}