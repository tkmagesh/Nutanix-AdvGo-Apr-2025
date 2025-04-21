package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100, 5)
	for primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int, workerCount int) <-chan int {
	primeCh := make(chan int)
	dataCh := produceData(start, end)
	go func() {
		workerWg := &sync.WaitGroup{}

		for range workerCount {
			workerWg.Add(1)
			go worker(dataCh, primeCh, workerWg)
		}
		workerWg.Wait()
		close(primeCh)
	}()
	return primeCh
}

func produceData(start, end int) <-chan int {
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func worker(dataCh <-chan int, primeCh chan<- int, workerWg *sync.WaitGroup) {
	for no := range dataCh {
		// fmt.Printf("worker : %d, no : %d\n", workerId+1, no)
		if isPrime(no) {
			primeCh <- no
		}
	}
	workerWg.Done()
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
