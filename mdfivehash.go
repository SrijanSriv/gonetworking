/* An md5 hash is a generic type of hash procedure in Go. Hashes are basically a writeup to some secure information.
Both the client and the server can have this writeup, and they can match that the output provided by taking the info
through this hash is the same or not. */
package main

import (
	"crypto/md5"
	"fmt"
)

func HashFunc(word string) {

	hash := md5.New()
	bytes := []byte(word)

	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
