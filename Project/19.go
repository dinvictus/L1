package main

import (
	"fmt"
	"unicode/utf8"
)

func Work19() {
	var str string = "главрыба"
	fmt.Println(str)
	str = Reverse_1(str)
	fmt.Println(str)
	str = Reverse_2(str)
	fmt.Println(str)
	str = Reverse_3(str)
	fmt.Println(str)
}

func Reverse_1(str string) string { // 186.6 ns/op		0 B/op		0 allocs/op // Не нужно дополнительной памяти
	runeStr := []rune(str)
	for left, right := 0, len(runeStr)-1; left < right; left, right = left+1, right-1 { // Меняем элементы местами двигаясь с разных сторон
		runeStr[left], runeStr[right] = runeStr[right], runeStr[left]
	}
	return string(runeStr)
}

func Reverse_2(str string) string { // 258.9 ns/op		64 B/op		1 allocs/op
	l := len(str)
	runeStr := make([]rune, l)
	l--
	for _, r := range str { // Проходим по строке и записываем руны в срез в обратном порядке
		runeStr[l] = r
		l--
	}
	return string(runeStr)
}

func Reverse_3(str string) string { //  150.4 ns/op		32 B/op		2 allocs/op // Самый быстрый
	size := len(str)             // Размер строки в байтах
	bufStr := make([]byte, size) // Используем байтовый срез
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(str[start:]) // Возвращает первую руну и её размер в байтах из входной строки
		start += n                                   // Прибавляем размер выданной руны к текущему размерру
		utf8.EncodeRune(bufStr[size-start:], r)      // Запись в байтовый срез с конца байтов, полученных из рун
	}
	return string(bufStr)
}
