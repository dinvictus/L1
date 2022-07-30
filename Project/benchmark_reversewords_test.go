package main

import "testing"

func BenchmarkReverseWords1(b *testing.B) {
	s := "snow dog sun hello World ! my name is dmitry"
	for i := 0; i < b.N; i++ {
		ReverseWords_1(s)
	}
}

func BenchmarkReverseWords2(b *testing.B) {
	s := "snow dog sun hello World ! my name is dmitry"
	for i := 0; i < b.N; i++ {
		ReverseWords_2(s)
	}
}

func BenchmarkReverseWords3(b *testing.B) {
	s := "snow dog sun hello World ! my name is dmitry"
	for i := 0; i < b.N; i++ {
		ReverseWords_3(s)
	}
}
