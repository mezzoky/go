package main

import "fmt"

var print = fmt.Println

func adder() func(int) int {
    sum := 0
    z := func (x int ) int {
        sum += x
        return sum
    }
    return z
}

func z() {

}

func y() func() {
    return z
}

func main() {
    var posz func(int) int = adder()
    print(posz(11), posz(11))

    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        print(
            pos(i),
            neg(-2*i),
        )
    }
}
