package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	var count int
	flag.IntVar(&count, "count", 0, "Number of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Spinning up %d goroutines.. Hit ENTER to start\n", count)
	fmt.Scanln()
	for idx := range count {
		wg.Add(1)        // increment the counter by 1
		go fn(idx+1, wg) // schedule the execution of this function through the scheduler
	}
	wg.Wait() // block until the counter becomes 0 (default = 0)
	fmt.Println("Done")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //simulate a time consuming operation
	fmt.Printf("fn[%d] completed\n", id)

}
