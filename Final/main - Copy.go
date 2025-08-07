package main

import (
	"bufio"
	"fmt"

	"log"
	"net"
)

func main() {
	connection, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	for {
		conn, err := connection.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
}
