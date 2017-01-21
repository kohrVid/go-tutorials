package main

import (
	//"runtime"
	"time"
)

func main() {
	//My laptop only has one CPU so this won't work
	//runtime.GOMAXPROCS(8)
	go abcGen()
	println("This comes first")
	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
