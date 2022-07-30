package main

import (
	"fmt"
)

func Work16() {
	Arr := []int{44, 14, 2, 4, 66, 54, 72, 34, 1, 3, 5, 8, 99, 12}
	// Quicksort(Arr)
	BubbleSort(Arr)
	fmt.Println(Arr)
}

func Quicksort(arr []int) { // 161.0 ns/op	0 B/op	0 allocs/op
	if len(arr) < 2 { // Если длина среза меньше двух, то заканчиваем сортировку
		return
	}
	slice := partition(arr) // Вызываем функцию, которая помещает числа меньше опорного элемента слева от него, а больше - справа
	Quicksort(arr[:slice])  // Делим срез на две части и рекурсивно вызываем функцию сортировки для каждой из частей
	Quicksort(arr[slice:])
}

func partition(arr []int) int {
	mid := arr[len(arr)/2] // Опорный элемент
	left := 0
	right := len(arr) - 1
	for {
		for arr[left] < mid { // Двигаемся слева пока не найдем элемент меньше опорного
			left++
		}
		for arr[right] > mid { // Двигаемся справа пока не найдём элемент больше опорного
			right--
		}
		if left >= right { // Если границы встретились, то заканчиваем и возвращаем пересечение
			return right
		}
		arr[left], arr[right] = arr[right], arr[left] // Меняем элементы местами
	}
}

// Сортировка пузырьком для сравнения эффективности
func BubbleSort(arr []int) { // 279.3 ns/op             0 B/op          0 allocs/op
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
