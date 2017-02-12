package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count uint32 = 0

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&count, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("count", atomic.LoadUint32(&count))
}
