package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100)
	for primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) <-chan int {
	primeCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go isPrime(no, primeCh, wg)
		}
		wg.Wait()
		close(primeCh)
	}()
	return primeCh
}

func isPrime(no int, primeCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	primeCh <- no
}
