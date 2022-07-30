package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Work8() {
	rand.Seed(time.Now().UnixNano())
	var numInt64 int64 = rand.Int63() // Получаем случайное число int64
	fmt.Println(numInt64)
	strNum2 := strconv.FormatInt(numInt64, 2) // Конвертируем int64 в двоичный формат и записываем в строку
	fmt.Println(strNum2)
	runeNum := []rune(strNum2) // Конвертируем строку в срез рун
	fmt.Println("Count bits: ", len(runeNum))
	var i int
	fmt.Scan(&i)
	if i < len(runeNum) && runeNum[i] == '0' { // Меняем i-й элемент(бит числа) руны на 1 или 0
		runeNum[i] = '1'
	} else if i < len(runeNum) && runeNum[i] == '1' {
		runeNum[i] = '0'
	}
	fmt.Println(string(runeNum))
	numInt64, _ = strconv.ParseInt(string(runeNum), 2, 64) // Конвертируем срез рун в строку и затем в int64
	fmt.Print(numInt64)
}
