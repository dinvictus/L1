package main

import "testing"

func BenchmarkSet2_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Work12_1()
	}
}

func BenchmarkSet2_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Work12_2()
	}
}
