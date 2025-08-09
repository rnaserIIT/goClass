package main

import (
	"crypto/tls"
	"io"
	"log"
)

func main() {
	// load client pem and key
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	// create a tls config to store the cert
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	// Use tls.Dial to connect securely
	conn, err := tls.Dial("tcp", "localhost:8000", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()

	log.Println("client: connected to server")

	// Send data to server
	_, err = io.WriteString(conn, "Hello Server\n")
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}

	log.Println("client: message sent")
}
