package utils

func IsPrime(no int64) bool {
	for i := int64(2); i <= no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(start, end int) []int {
	primes := make([]int, 0, 100)
	// var primes []int
	for no := start; no <= end; no++ {
		if IsPrime(int64(no)) {
			primes = append(primes, no)
		}
	}
	return primes
}
