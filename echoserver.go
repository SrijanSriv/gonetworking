/* The `go` before the function handleClient makes it multi-threaded. This basically means that multiple client can
connect to this function. This demonstrates the power of go; ease of multithreading */
package main

import "net"

func EchoServer(port string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	defer conn.Close()

	var buffer [512]byte

	for {
		n, err1 := conn.Read(buffer[0:])
		if err1 != nil {
			return
		}

		_, err2 := conn.Write(buffer[:n])
		if err2 != nil {
			return
		}
	}
}
