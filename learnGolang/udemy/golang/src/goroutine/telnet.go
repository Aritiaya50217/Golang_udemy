package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// client
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// copy conn --> Stdout
	//log.Println("copy from conn to stdout")

	log.Println("Connected to server successfully.")
	// goroutine
	go io.Copy(os.Stdout, conn)
	//log.Println("copy from stdin to conn")
	io.Copy(conn, os.Stdin)
}
