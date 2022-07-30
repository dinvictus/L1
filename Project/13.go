package main

import "fmt"

func Work13() {
	var a int = 10
	var b int = 40
	fmt.Println("a:", a, "b:", b)
	// Не сработает, если сумма a + b больше, чем может хранить int
	a = a + b // a = 50
	b = a - b // b = 50 - 40 = 10
	a = a - b // a = 50 - 10 = 40
	fmt.Println("a:", a, "b:", b)
	// Используем XOR, манипулирует с битами
	a = a ^ b // 001010 ^ 101000 = 100010
	b = a ^ b // 100010 ^ 101000 = 001010
	a = a ^ b // 100010 ^ 001010 = 101000
	fmt.Println("a:", a, "b:", b)
	// Используем встроенные функции языка
	a, b = b, a
	fmt.Println("a:", a, "b:", b)
}
