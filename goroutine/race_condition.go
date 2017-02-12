package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var total int32 = 10
var mu = &sync.Mutex{}

func sell(id int) {
	for {
		mu.Lock()
		if total > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			total--
			fmt.Println("id:", id, "ticket:", total)
		} else {
			break
		}
		mu.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		go sell(i)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println(total, "done")
}
