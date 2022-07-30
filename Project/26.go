package main

import (
	"fmt"
	"unicode"
)

func Work26() {
	str := "aBcDeF"
	check := CheckStr_1(str)
	fmt.Println(check)
	str2 := "abCdefAf"
	check2 := CheckStr_1(str2)
	fmt.Println(check2)
}

func CheckStr_1(str string) bool { // 601.9 ns/op		96 B/op		1 allocs/op
	MapForSymbols := make(map[rune]struct{}) // Map для хранения символов в ключах
	for _, r := range str {                  // Проходим по строке
		if _, ok := MapForSymbols[r]; !ok { // Если в map уже есть такой ключ, то возвращаем false
			MapForSymbols[unicode.ToLower(r)] = struct{}{}
		} else {
			return false
		}
	}
	return true // Если одинаковых ключей не найдено, то возвращаем true
}

func CheckStr_2(str string) bool { // 1189 ns/op		0 B/op		0 allocs/op
	for i, r1 := range str {
		for j, r2 := range str {
			if unicode.ToLower(r1) == unicode.ToLower(r2) && i != j {
				return false
			}
		}
	}
	return true
}
