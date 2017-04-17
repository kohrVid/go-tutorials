// S07a: package io contains many useful readers.
//
// OUTPUT:
//
//    $ cat hello.txt | go run main.go
//    Reflection is never clear.
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// TODO: Only read the first 27 bytes from standard input (3/6 lines).
	if _, err := io.CopyN(os.Stdout, os.Stdin, 27); err != nil {
		log.Fatal(err)
	}
}
