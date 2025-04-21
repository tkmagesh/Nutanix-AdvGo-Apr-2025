package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of this function through the scheduler
	f2()
	time.Sleep(2 * time.Second) // block the execution of main so that the scheduler can look for other goroutines scheduled and execute them (cooperative multitasking)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second) //simulate a time consuming operation
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
