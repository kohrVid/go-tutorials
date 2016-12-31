/*
  From http://ridiculousfish.com/blog/posts/go_bloviations.html
*/
package main

import (
	"fmt"
)

func main() {
	values := []string{"a", "b", "c"}
	for _, v := range values {
		var done chan bool
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
}
