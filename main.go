/* The driver function. All the functions can be tested from here. To connect to a `server` i.e. a function that is
`listening` to a certain connection locally, we may use, in a seperate terminal from the one running the server itself,
the following command for tcp connections:
`telnet localhost "port"`
or for udp conncections we may use:
`netcat localhost "port" -u` followed by a random string for a message.
Here, the port is the one provided when setting up the server. It is usually :1200 */
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

	HashFunc(input)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
