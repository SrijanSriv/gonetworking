package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func TLSEchoServer(port string) {

	service := port
	cert, err := tls.LoadX509KeyPair("somerandomfile.pem", "private.pem") // both of these are a PEM file, that i cant supply
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", service, &config)
	checkError(err)

	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClientTLS(conn)
	}
}

func handleClientTLS(conn net.Conn) {

	defer conn.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
