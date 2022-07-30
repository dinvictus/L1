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

func ReverseWords_1(str string) string { // 995.4 ns/op		456 B/op		11 allocs/op
	splitStr := strings.Split(str, " ") // Разбиваем строку на срез строк по пробелу
	resStr := ""
	for i := len(splitStr) - 1; i >= 1; i-- { // Проходим по срезу в обратном порядке
		resStr += splitStr[i] + " " // Записываем слова в строку
	}
	resStr += splitStr[0]
	return resStr
}

func ReverseWords_2(str string) string { // 508.4 ns/op		272 B/op		3 allocs/op  // Самый быстрый
	splitStr := strings.Split(str, " ")
	var buffer bytes.Buffer // Используем байтовый буфер для записи слов
	for i := len(splitStr) - 1; i >= 1; i-- {
		buffer.WriteString(splitStr[i])
		buffer.WriteString(" ")
	}
	buffer.WriteString(splitStr[0])
	return buffer.String()
}

func ReverseWords_3(str string) string { // 544.6 ns/op		280 B/op		5 allocs/op
	splitStr := strings.Split(str, " ")
	var strBuilder strings.Builder // Используем string.Builder для записи слов
	for i := len(splitStr) - 1; i >= 1; i-- {
		strBuilder.WriteString(splitStr[i])
		strBuilder.WriteString(" ")
	}
	strBuilder.WriteString(splitStr[0])
	return strBuilder.String()
}
