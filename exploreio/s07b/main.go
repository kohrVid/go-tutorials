// S07b: Read sections from a reader.
//
// OUTPUT:
//
//     $ go run main.go
//     io.Reader
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader to be read\n")

	// TODO: Print the string "io.Reader" to stdout (4 lines).
	section := io.NewSectionReader(r, 5, 9)
	if _, err := io.Copy(os.Stdout, section); err != nil {
		log.Fatal(err)
	}
	fmt.Println("") // for readability
}
