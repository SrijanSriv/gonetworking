/* Just a simple function to notice how a certain server reacts when its header is asked about */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func GetHeadInfo(service string) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}
