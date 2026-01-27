package main

import (
	"fmt"
	"sync"
	"time"
)

// 1- WaitGroup
func basicWaitGroupExample() {
	var wg sync.WaitGroup

	numberOfRoutines := 7
	wg.Add(numberOfRoutines)
	// launches 7 goroutines

	for i := 1; i <= numberOfRoutines; i++ {
		go basicWorker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

func basicWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // call done via defer
	fmt.Printf("Goroutine %d is running\n", id)
}


// 2 - WaitGroup + sleep
func sleepExample() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sleepWorker(1, &wg)
	go sleepWorker(2, &wg)
	go sleepWorker(3, &wg)

	wg.Wait()
	fmt.Println("Main finished")
}

func sleepWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// 3 - nested goroutines
func nestedExample() {
	var wg sync.WaitGroup

	wg.Add(2)

	go parentWorker(1, &wg)
	go parentWorker(2, &wg)

	wg.Wait()
	fmt.Printf("Main routine (nested)\n")
}

func parentWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Parent %d running\n",  id)

// inner WaitGroup for child goroutine
	var innerWg sync.WaitGroup
	innerWg.Add(1)

// child goroutine (anonymous func)
	go func() {
		defer innerWg.Done()
		fmt.Printf("Child of %d is running\n", id)
	}()

	innerWg.Wait()
	fmt.Printf("Parent %d finished\n", id)
}


func main() {
	basicWaitGroupExample()
	// sleepExample()
	// nestedExample()

}