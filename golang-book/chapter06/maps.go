package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Printf("\"key\" => %d\n", x["key"])
}
