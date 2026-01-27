package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait until ALL goroutines finish their work
	wg := &sync.WaitGroup{}
	mut:= &sync.Mutex{}
	// mutual exlcusion lock 
	var score = []int{0}
	
	// no of workers
	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		// create a new slice and return it
		// then assign it back to score
		score = append(score,1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// passing wg variable when starting the goroutine

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score,2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score,3)
		mut.Unlock()
		wg.Done()
	} (wg, mut)

	// wait for all goroutines 
	wg.Wait()	
	fmt.Println(score)
}
