package main

import (
	"testing"
)

func BenchmarkMultiValueAssignment(b *testing.B) {
	a, x, c := 1, 2, 3
	for i := 0; i < b.N; i++ {
		a, x, c = a+x, a+x, x
	}
	_ = a
	_ = b
	_ = c
}

func BenchmarkSeparateAssignment(b *testing.B) {
	a, x, c := 1, 2, 3
	for i := 0; i < b.N; i++ {
		a = a + x
		x = a + x
		c = x
	}
	_ = a
	_ = b
	_ = c
}
