package benchmark_go

import (
	"testing"
)

func fibRecursive(n uint) uint {
	if n <= 2 {
		return 1
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(position uint) uint {
	slc := make([]uint, position)
	slc[0] = 1
	slc[1] = 1

	if position <= 2 {
		return 1
	}
	var result, i uint
	for i = 2; i < position; i++ {
		result = slc[i-1] + slc[i-2]
	}
	return result
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibRecursive(uint(60))
	}
}

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibIterative(uint(60))
	}
}
