package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genData()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genData() <-chan int {
	ch := make(chan int)

	go func() {
		count := rand.Intn(20)
		fmt.Printf("[genData] producing %d numbers\n", count)
		for i := range count {
			ch <- (i + 1) * 10
		}
		close(ch)
	}()
	return ch
}
