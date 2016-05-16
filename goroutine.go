package main
import (
    "fmt"
)
var print = fmt.Println
var scan = fmt.Scanln

const UNTIL = 5555

var t1done = make(chan bool)
var t2done chan bool  = make(chan bool)


func main() {
    print("start")
    go task1()
    go task2()
    <- t1done
    done2 := <- t2done

    go runner("goroutine")
    go func(msg string) {
        print(msg)
    }("going")


    var input string
    scan(&input)

    print("end", done2)
}

func task1() {
    print("task1 start")
    for i := 0; i < UNTIL; i++ {
        print("task1", i)
    }
    print("task1 end")
    t1done <- true
}

func task2() {
    print("task2 start")
    for i := 0; i < UNTIL; i++ {
        print("task2", i)
    }
    print("task2 end")
    t2done <- true
}

func runner(from string) {
    for i := 0; i < 5; i++ {
        print(from, ":", i)
    }
}
