package main

import "fmt"

func Work10() {
	MapTemperature := make(map[int][]float32) // map для группировки значений температуры
	var arr []float32 = []float32{-25.4, -27.0, 13.0, 19.0, -11.1, -12.2, -19.9, 15.5, 24.5, -21.0, 32.5, -20.0, -30.0, 0.0, 0.1, 10.1, 10.0, 20.0, 30.0, -5.5, -1.2, -7.8, 9.9, 3.2, 8.8}
	for _, el := range arr {
		key := (int(el) / 10) * 10                            // Вычисление ключа для map
		MapTemperature[key] = append(MapTemperature[key], el) // Добавление значения в промежуток по ключу
	}
	fmt.Println(MapTemperature)
}
