package main

import (
	"fmt"
	"math/big"
)

func Work22() {
	a := new(big.Int) // Используем пакет math big для работы с большими числами
	b := new(big.Int)
	a.SetString("10000000000000000", 10) // Устанавливаем значения из строки
	b.SetString("5000000000000000", 10)
	c := new(big.Int)
	c.Div(a, b) // Деление
	fmt.Println(c)
	b.Mul(c, a) // Умножение
	fmt.Println(b)
	a.Add(b, c) // Сумма
	fmt.Println(a)
	c.Sub(b, a) // Разность
	fmt.Println(c)
}
