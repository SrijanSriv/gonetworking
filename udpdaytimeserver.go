/* A daytime server but built in UDP over IP. The other one was TCP over IP. */
package main

import (
	"net"
	"time"
)

func UDPDaytimeServer(port string) {

	var conn *net.UDPConn

	udpAddr, err := net.ResolveUDPAddr("udp", port)
	checkError(err)

	conn, err = net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClientUDP(conn)
	}
}

func handleClientUDP(conn *net.UDPConn) {
	var buffer [512]byte

	_, addr, err := conn.ReadFromUDP(buffer[0:])
	if err != nil {
		return
	}

	daytime := time.Now().Format("02-01-2006 15:04:05")

	conn.WriteToUDP([]byte(daytime), addr)
}
