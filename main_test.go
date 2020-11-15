package main

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for _, size := range []int{1, 2, 3, 10, 20, 40, 60} {
		benchmarkFib(b, size)
	}
}

var result int

func benchmarkFib(b *testing.B, size int) {
	b.Run(fmt.Sprintf("Fib_%d", size), func(b *testing.B) {
		var r int
		for i := 0; i < b.N; i++ {
			r = Fib(size)
		}

		result = r
	})
}
