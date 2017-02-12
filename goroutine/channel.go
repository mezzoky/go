package main

// http://coolshell.cn/articles/8489.html

import (
	"fmt"
	"time"
)

func basic() {
	var channel = make(chan string, 1)

	go func() {
		channel <- "hi"
		channel <- "world"
	}()

	var m1 = <-channel
	var m2 = <-channel

	fmt.Println(m1, m2)
}

func readWrite() {
	// channel by default is blocking:
	//	block at writer when full, block at reader when empty
	// channel buffer default is 1, which will block "world" in writer
	// if "hello" haven't received by msg
	// if the channel buffer is more than 1, "world" will run after "hello" even
	// the receiver / reader hvnt receive any msg
	var channel = make(chan string, 1)

	go func() {
		channel <- "hello"
		fmt.Println("write hello")

		channel <- "world"
		fmt.Println("write world")

		fmt.Println("sleep 3 seconds")
		time.Sleep(3 * time.Second)

		channel <- "channel"
		fmt.Println("write channel")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("read wake up")

	var msg = <-channel
	fmt.Println("read", msg)

	msg = <-channel
	fmt.Println("read", msg)

	msg = <-channel
	fmt.Println("read", msg)
}

func multiChannelBlockingSelect() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("receive m1", m1)
		case m2 := <-c2:
			fmt.Println("receive m2", m2)
		}
	}
}

func multiChannelBlockingTimeout() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2"
	}()

	var timeout = 0
	for {
		select {
		case m1 := <-c1:
			fmt.Println("receive m1", m1)
		case m2 := <-c2:
			fmt.Println("receive m2", m2)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			timeout++
		}

		if timeout > 3 {
			break
		}
	}
}

func multiChannelNonBlocking() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2"
	}()

	for {
		select {
		case m1 := <-c1:
			fmt.Println("receive m1", m1)
		case m2 := <-c2:
			fmt.Println("receive m2", m2)
		default:
			fmt.Println("nothing received")
			time.Sleep(2 * time.Second)
		}
	}
}

func channelClose() {

}

func main() {
	multiChannelNonBlocking()
}
