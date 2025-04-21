package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) []int {
	result := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			result = append(result, no)
		}
	}
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
