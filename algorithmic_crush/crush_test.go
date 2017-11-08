package main

import "testing"

func BenchmarkValidate10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddElements2(1, 10000000, 100)
	}
}
