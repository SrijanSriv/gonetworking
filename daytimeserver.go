/* A daytime server can be connected to just to get the daytime, after which it is expected that the connection is closed */
package main

import (
	"net"
	"time"
)

func DaytimeServer(port string) {

	var conn net.Conn

	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err = listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().Format("02-01-2006 15:04:05")
		daytime += "\nBrought to you by srijansriv\n"
		conn.Write([]byte(daytime))
		conn.Close()
	}

}
