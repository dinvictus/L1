package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Work20() {
	str := "snow dog sun hello World ! my name is dmitry привет как дела"
	fmt.Println(str)
	str = ReverseWords_1(str)
	fmt.Println(str)
	str = ReverseWords_2(str)
	fmt.Println(str)
	str = ReverseWords_3(str)
	fmt.Println(str)
}

func ReverseWords_1(str string) string { // 34349 ns/op		56248 B/op		151 allocs/op
	splitStr := strings.Split(str, " ") // Разбиваем строку на срез строк по пробелу
	resStr := ""
	for i := len(splitStr) - 1; i >= 1; i-- { // Проходим по срезу в обратном порядке
		resStr += splitStr[i] + " " // Записываем слова в строку
	}
	resStr += splitStr[0]
	return resStr
}

func ReverseWords_2(str string) string { // 6732 ns/op		5616 B/op		7 allocs/op
	splitStr := strings.Split(str, " ")
	var buffer bytes.Buffer // Используем байтовый буфер для записи слов
	for i := len(splitStr) - 1; i >= 1; i-- {
		buffer.WriteString(splitStr[i])
		buffer.WriteString(" ")
	}
	buffer.WriteString(splitStr[0])
	return buffer.String()
}

func ReverseWords_3(str string) string { // 5525 ns/op		4600 B/op		9 allocs/op
	splitStr := strings.Split(str, " ")
	var strBuilder strings.Builder // Используем string.Builder для записи слов
	for i := len(splitStr) - 1; i >= 1; i-- {
		strBuilder.WriteString(splitStr[i])
		strBuilder.WriteString(" ")
	}
	strBuilder.WriteString(splitStr[0])
	return strBuilder.String()
}
