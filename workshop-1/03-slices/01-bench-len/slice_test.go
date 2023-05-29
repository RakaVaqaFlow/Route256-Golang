package main

import (
	"fmt"
	"testing"
)

func FillSlice() []int {
	const sz = 32000
	a := make([]int, 0, sz)
	for i := 0; i < sz; i++ {
		a = append(a, i)
	}
	fmt.Print(a)
	return a
}

func BenchmarkFillSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FillSlice()
	}
}
