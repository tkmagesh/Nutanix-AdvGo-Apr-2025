package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := time.After(5 * time.Second)
	ch := genData(stopCh)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

	fmt.Println("Done")
}

func genData(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)

	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received")
				break LOOP
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}
