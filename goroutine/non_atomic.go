package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var count uint32 = 0

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				count++
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("count", count)
}
