package main

import . "fmt"

func main() {
	print("mew", "two")
}

func print(v ...interface{}) {
	Println(v...)
}
