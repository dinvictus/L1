package main

import "testing"

func BenchmarkSortQuick(b *testing.B) {
	Arr := []int{44, 14, 2, 4, 66, 54, 72, 34, 1, 3, 5, 8, 99, 12}
	for i := 0; i < b.N; i++ {
		Quicksort(Arr)
	}
}

func BenchmarkSortBubble(b *testing.B) {
	Arr := []int{44, 14, 2, 4, 66, 54, 72, 34, 1, 3, 5, 8, 99, 12}
	for i := 0; i < b.N; i++ {
		BubbleSort(Arr)
	}
}
