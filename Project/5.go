package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Work5() {
	channel := make(chan int) // Канал для передачи данных
	wg := new(sync.WaitGroup) // WaitGroup для ожидания завершения воркера
	wg.Add(1)                 // Один воркер
	go Worker5(channel, wg)   // Запуск воркера, передаём канал и WaitGroup
	var seconds int
	fmt.Scan(&seconds)
	stop := time.After(time.Second * time.Duration(seconds)) // Канал, который через seconds секунд вернёт значение
	for loop := true; loop; {                                // Цикл для отправки данных в канал
		select {
		case <-stop: // Когда канал stop вернёт значение, выполнится этот блок, выходим из цикла
			loop = false
		default: // В ином случае выполится блок default
			randInt := rand.Int() // Генерация случайного числа int
			channel <- randInt    // Отправка числа в канал
			fmt.Print("Send: ", randInt)
			time.Sleep(time.Millisecond * 500)
		}
	}
	fmt.Println("Time is over")
	close(channel) // Закрытие канала, цикл в воркере завершается
	wg.Wait()      // Ждем завершения работы воркера
	fmt.Println("End program")
}

func Worker5(channel chan int, wg *sync.WaitGroup) { // Воркер
	for number := range channel { // Цикл чтения данных из канала
		fmt.Print(" Recieve: ", number, "\n")
	}
	wg.Done() // Уменьшаем значение WaitGroup на 1, завершение работы воркера
	fmt.Println("Worker is done")
}
