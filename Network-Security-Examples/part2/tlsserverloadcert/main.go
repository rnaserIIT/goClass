package main

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	// Load TLS certificate and key
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatal("server: loadkeys: ", err)
	}

	// Set up TLS config
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		Rand:         rand.Reader,
	}

	// Start TLS listener
	listener, err := tls.Listen("tcp", ":8000", config)
	if err != nil {
		log.Fatal("server: listen: ", err)
	}
	defer listener.Close()

	log.Println("server: listening on :8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("server: accept error:", err)
			continue
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
