/* A certain dns may have multiple `aliases`. All these can be found out by calling the following function */
package main

import (
	"fmt"
	"net"
	"os"
)

func DnsLookupMulti(name string) {

	/* net.LookupCNAME(name) can be used to look up canonical host name*/
	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
