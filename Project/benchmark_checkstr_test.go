package main

import "testing"

func BenchmarkCheckStr1(b *testing.B) {
	str := "AbCdEfGhigklmnopqrstuvwxyZz"
	for i := 0; i < b.N; i++ {
		CheckStr_1(str)
	}
}

func BenchmarkCheckStr2(b *testing.B) {
	str := "AbCdEfGhigklmnopqrstuvwxyZz"
	for i := 0; i < b.N; i++ {
		CheckStr_2(str)
	}
}
