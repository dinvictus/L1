package main

import (
	"fmt"
)

func Work17() {
	Arr := []int{44, 14, 2, 4, 66, 54, 72, 34, 1, 3, 5, 8, 99, 12}
	Quicksort(Arr) // [1 2 3 4 5 8 12 14 34 44 54 66 72 99]
	fmt.Println(Arr)
	index := BinSearch(Arr, 66)
	fmt.Println(index)
}

func BinSearch(arr []int, num int) int { // 28.03 ns/op		0 B/op		0 allocs/op
	return part(arr, num, 0, len(arr)) // Изначально вызываем функцию part с промежутком от начала и до конца массива
}

func part(arr []int, num, left, right int) int {
	midIndex := left + (right-left)/2 // Находим средний индекс в заданном промежутке массива
	// fmt.Println(left, " ", right, " ", midIndex)
	if right-left <= 1 && arr[midIndex] != num { // Если длина промежутка меньше либо равна 1 и элемент с средним индексом не равен искомому, то элемента в массиве нет
		return -1
	}
	if num > arr[midIndex] { // Если искомый элемент больше среднего, то рекурсивно вызываем эту функцию с промежутком зачений, находяшихся справа от среднего
		midIndex = part(arr, num, midIndex, right)
	} else if num < arr[midIndex] { // Если искомый элемент меньше среднего, то рекурсивно вызываем эту функцию с промежутком зачений, находяшихся слева от среднего
		midIndex = part(arr, num, left, midIndex)
	}
	return midIndex // Если ни одно из условий не сработало, то искомый элемент найден, возвращаем его индекс
}

// Обычный поиск для сравнения эффективности
func SimpleSearch(arr []int, num int) int { // 652.3 ns/op		0 B/op		0 allocs/op
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
