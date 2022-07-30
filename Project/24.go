package main

import (
	"fmt"
	p "point"
)

func Work24() {
	p1 := p.Point(10, 12) // Создаём объекты структуры point с помощью функции-конструктора
	p2 := p.Point(15, 18)
	distance := p1.Distance(p2) // Вызываем метод для нахождения дистанции
	fmt.Println(distance)
}
