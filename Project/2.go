package main

import (
	"fmt"
	"sync"
)

func Work2() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := new(sync.WaitGroup) // WaitGroup для того, чтобы дождаться выполнения всех горутин
	wg.Add(10)                // Значение 10, так как размер массива 10
	for _, el := range arr {
		go func(el int) { // Запуск анонимной горутины на каждой итерации цикла
			fmt.Print(el*el, " ") // Вывод квадрата числа
			wg.Done()             // Уменьшение значения WaitGroup на 1, конец работы текущей горутины
		}(el)
	}
	wg.Wait() // Ожидание завершения всех горутин
}
