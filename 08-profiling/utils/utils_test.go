package utils_test

import (
	"profiling-demo/utils"
	"testing"
)

func Benchmark_GeneratePrimes(b *testing.B) {
	for range b.N {
		utils.GeneratePrimes(1000, 2000)
	}
}
