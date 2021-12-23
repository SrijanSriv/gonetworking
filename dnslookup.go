/* Simple function that looks up for the ip address of a specific domain */
package main

import (
	"fmt"
	"net"
	"os"
)

func DnsLookup(dns string) {

	addr, err := net.ResolveIPAddr("ip", dns)

	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resloved address is:", addr.String())
	os.Exit(0)

}
