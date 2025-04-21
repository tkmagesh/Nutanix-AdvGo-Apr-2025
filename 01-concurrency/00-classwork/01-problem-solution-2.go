package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) []int {
	result := make([]int, 0)
	var mutex sync.Mutex
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				mutex.Lock()
				{
					result = append(result, no)
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
