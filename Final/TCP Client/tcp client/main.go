package main

import (
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(conn, "Hello Server")
	conn.Close()
}
