package main

import "crypto/tls"

func TLSEchoClient(port string) {

	_, err := tls.Dial("tcp", port, nil) // _ should be conn. figuring out how to create a more universal client.
	checkError(err)

}
