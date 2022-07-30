package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Work4() {
	channel := make(chan int)                      // Канал для передачи данных в горутины
	channelShutdown := make(chan os.Signal, 1)     // Канал для перехвата сигнала
	signal.Notify(channelShutdown, syscall.SIGINT) // Прослушивание и запись сигнала в канал
	wg := new(sync.WaitGroup)                      // WaitGroup для ожидания окончания работы всех горутин
	var countWorkers int                           // Количество воркеров
	fmt.Scan(&countWorkers)
	wg.Add(countWorkers)                // Установка количества воркеров для WaitGroup
	for i := 0; i < countWorkers; i++ { // Запуск воркеров
		go Worker4(channel, i, wg) // Передаём канал данных, порядковый номер и WaitGroup
	}
	for loop := true; loop; { // Цикл для отправки данных
		select {
		case <-channelShutdown: // При получении сигнала в канал, выполнится этот блок, выходим из цикла
			loop = false
		default: // В ином случае выполняется этот блок
			channel <- rand.Int() // Отправка в канал случайного числа
			time.Sleep(10 * time.Millisecond)
		}
	}
	close(channel) // Закрываем канал, все циклы в воркерах завершаются
	close(channelShutdown)
	wg.Wait() // Ожидаем завершения работы всех воркеров
	fmt.Println("All workers is done")
}

func Worker4(channel chan int, number int, wg *sync.WaitGroup) { // Воркер
	for num := range channel { // Чтение данных из канала
		fmt.Println("Hello, I'm a", number, "Worker. My number: ", num) // Вывод данных
	}
	wg.Done() // Уменьшаем значение WaitGroup на 1, завершение работы воркера
	fmt.Println("Worker", number, "is done.")
}
