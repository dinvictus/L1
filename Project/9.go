package main

import (
	"fmt"
	"sync"
)

func Work9() {
	var arr [20]int = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	channel1 := make(chan int)           // Первый канал для отправки и получения обычных чисел x
	channel2 := make(chan int)           // Второй канал для получения и вывода чисел x*2
	go Worker9_1(channel1, channel2, wg) // Запуск первого воркера, передаём два канала
	go Worker9_2(channel2, wg)           // Запуск второго воркера, передаём только второй канал для вывода
	for _, x := range arr {
		channel1 <- x // Отправка x в канал
	}
	close(channel1) // Закрытие первого канала заканчивает цикл чтения в первом воркере
	wg.Wait()       // Ожидание завершения всех воркеров
}

func Worker9_1(channel1, channel2 chan int, wg *sync.WaitGroup) {
	for x := range channel1 { // Чтение данных из первого канала
		channel2 <- x * 2 // Отправка во второй канал значений x * 2
	}
	close(channel2) // После закрытия первого канала и выхода из цикла чтения, закрываем второй канал, что закончит цикл чтения во втором воркере
	wg.Done()       // Завершение работы первого воркера
}

func Worker9_2(channel chan int, wg *sync.WaitGroup) {
	for num := range channel { // Чтение данных из второго канала
		fmt.Println(num) // Вывод данных из второго канала
	}
	wg.Done() // Завершение работы второго воркера
}
