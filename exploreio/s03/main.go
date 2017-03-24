// S03: Stdin is a file, too.
//
// A filter that does nothing.
//
// OUTPUT:
//
//     $ cat hello.txt | go run main.go
//     Cgo is not Go.
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// TODO: Read input from standard input and pass it to standard output,
	//       without using a byte slice (3 lines).
	// ...
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
	// ...
}
