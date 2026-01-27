package main

import ("fmt";
 "time")

func main() {
	go fmt.Printf("New routine ")
	// adding a delay is bad! time assumptions may be wrong 
	time.Sleep(1 * time.Millisecond)
	// schedulers favors the main routine
	fmt.Printf("Main routine")

}  