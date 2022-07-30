package main

import "testing"

func BenchmarkSet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Work11_1()
	}
}

func BenchmarkSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Work11_2()
	}
}
