package main

import (
	"fmt"
	"runtime"
	//"sync"
)

func main() {
	//mutex := new(sync.Mutex)
	runtime.GOMAXPROCS(4)
	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			//mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
				//mutex.Unlock()
			}()
			fmt.Scanln()
		}
	}
}
