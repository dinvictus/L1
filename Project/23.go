package main

import "fmt"

func Work23() {
	Arr := []int{44, 14, 2, 4, 66, 54, 72, 34, 1, 3, 5, 8, 99, 12}
	Arr = Del_1(Arr, 4)
	fmt.Println(Arr)
	Arr = Del_2(Arr, 7)
	fmt.Println(Arr)
}

func Del_1(arr []int, index int) []int { // Изменяет порядок, сложность постоянная
	arr[index] = arr[len(arr)-1] // Меняем последний и удаляемый элемент местами
	arr[len(arr)-1] = 0          // Обнуляем последний элемент
	return arr[:len(arr)-1]      // Возвращаем срез усечённый на 1
}

func Del_2(arr []int, index int) []int { // Сохраняет порядок, сложность зависит от размера среза
	copy(arr[index:], arr[index+1:]) // копируем срез от индекс + 1 до конца в срез от индекс до конца
	arr[len(arr)-1] = 0              // обнуляем последний элемент
	return arr[:len(arr)-1]          // Возвращаем срез усечённый на 1
}
