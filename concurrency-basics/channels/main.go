package main
import (
    "fmt"
    "time")

// worker func 
// sends result to channel
func prod(v1 int, v2 int, c chan int) {

    result := v1* v2
    // fmt.Println(result)
    c <- result

}

func main() {
    // create unbuffered channel 
    c:= make(chan int)

    // start multiple goroutines
    go prod(11,2,c)
    go prod(32,4,c)

    fmt.Println("main: sleeping before receiving")
	time.Sleep(1 * time.Second)

    fmt.Println("main: ready to receive")

    a:= <- c
    fmt.Println("main: received", a)

    b:= <- c
    fmt.Println("main: received", b)

    finalResult := a * b 
    fmt.Println("final result:", finalResult)
}