package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

func main() {
	// load client pem and key
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	// create a tls config to store the cert
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(conn, "Hello Server")
	conn.Close()
}
