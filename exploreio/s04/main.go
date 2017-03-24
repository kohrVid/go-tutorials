// S04: Decompression with gzip and a simple filter chain.
//
// OUTPUT:
//
//     $ cat hello.txt.gz | go run main.go
//     Don't just check errors, handle them gracefully.
package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {

	// TODO: Read gzip compressed data from standard input and write
	//       to standard output, without using a byte slice (7 lines).
	b, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, b); err != nil {
		log.Fatal(err)
	}
}
