/* The driver function. All the functions can be tested from here. To connect to a `server` i.e. a function that is
`listening` to a certain connection locally, we may use, in a seperate terminal from the one running the server itself,
the following command:
telnet localhost `port`
where the port is the one provided when setting up the server */
package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	input := os.Args[1]

	EchoServer(input)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
