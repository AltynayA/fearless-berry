package main

import (
	"sync"
	"fmt"
)

var on sync.Once
var wg sync.WaitGroup

func setup() {
	fmt.Println("Init")
}

func dostuff() {
	defer wg.Done() 
	// init
	on.Do(setup)
	fmt.Println("Hello")
}


func main() {
	num := 3
	wg.Add(num)

	for i := 0; i < num; i++ {
		go dostuff()
	}

	wg.Wait() 
}