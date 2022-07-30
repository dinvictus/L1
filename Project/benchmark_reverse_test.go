package main

import "testing"

func BenchmarkReverse1(b *testing.B) {
	s := "главрыба"
	for i := 0; i < b.N; i++ {
		Reverse_1(s)
	}
}

func BenchmarkReverse2(b *testing.B) {
	s := "главрыба"
	for i := 0; i < b.N; i++ {
		Reverse_2(s)
	}
}

func BenchmarkReverse3(b *testing.B) {
	s := "главрыба"
	for i := 0; i < b.N; i++ {
		Reverse_3(s)
	}
}
