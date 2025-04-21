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
	dataCh := make(chan int)
	primeCh := make(chan int)
	go func() {
		go func() {
			for no := start; no <= end; no++ {
				dataCh <- no
			}
			close(dataCh)
		}()

		workerCount := 5
		workerWg := &sync.WaitGroup{}
		// for workerId := range workerCount {
		for range workerCount {
			workerWg.Add(1)
			go func() {
				for no := range dataCh {
					// fmt.Printf("worker : %d, no : %d\n", workerId+1, no)
					if isPrime(no) {
						primeCh <- no
					}
				}
				workerWg.Done()
			}()
		}
		workerWg.Wait()
		close(primeCh)
	}()
	return primeCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
