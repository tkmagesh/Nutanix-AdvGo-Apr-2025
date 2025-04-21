package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the counter by 1
	go f1()   // schedule the execution of this function through the scheduler
	f2()

	// Poor man's synchronization technique (DO NOT USE THIS)
	// time.Sleep(2 * time.Second) // block the execution of main so that the scheduler can look for other goroutines scheduled and execute them (cooperative multitasking)

	wg.Wait() // block until the counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //simulate a time consuming operation
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
