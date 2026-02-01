package main

import (
	"fmt"
	"sync"
	"time" 
)

// structs

// host allows <= 2 philosophers to eat concurrently
type Host struct {
	channel chan struct{}
	// buffer channel to limit 
}

// chopstick is a mutex
type Chopstick struct {
	sync.Mutex
}


// philosopher 
type Philosopher struct {
	id              int
	leftCS, rightCS *Chopstick
}


// eat method
func (p *Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {

		// ask host for permission
		host.channel <- struct{}{}

		// pick up chopsticks (order differs per philosopher)
		if p.id%2 == 0 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		// start eating after lock
		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("finishing eating %d\n", p.id)

		// put down chopsticks
		p.leftCS.Unlock()
		p.rightCS.Unlock()

		// release permission
		<-host.channel
	}
}

func main() {
	var wg sync.WaitGroup
	// host with 2 capacity
	host:= Host{channel: make(chan struct{}, 2)}

	// create chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{}
	}
	// create philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:      i + 1,
			leftCS: chopsticks[i],
			rightCS: chopsticks[(i+1)%5],
		}
	}
	// add goroutines
	wg.Add(5)
	for i:=0; i <5; i++ {
		go philosophers[i].eat(&host, &wg)
	}

	wg.Wait()
}