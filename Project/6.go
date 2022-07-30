package main

import (
	"fmt"
	"sync"
	"time"
)

func Work6() {
	wgAll := new(sync.WaitGroup) // WaitGroup для ожидания завершения всех горутин
	wgAll.Add(5)
	// Горутина 1
	channelStopG1 := make(chan int)     // Канал для горутины 1
	go Goroutine1(channelStopG1, wgAll) // Запуск горутины 1
	time.Sleep(time.Second)
	channelStopG1 <- 0   // Передаём значение в канал, горутина 1 продолжает работу
	close(channelStopG1) // Такой же эффект будет иметь закрытие канала
	// Горутина 2
	channelStopG2 := make(chan int)     // Канал для горутины 2
	go Goroutine2(channelStopG2, wgAll) // Запуск горутины 2
	time.Sleep(time.Second)
	<-channelStopG2 // Горутина 2 будет заблокирована, пока значение из канала не будет прочитано, так как он не буфферизированный
	// Горутина 3
	mu := new(sync.Mutex)    // Используем мьютекс для остановки выполнения горутины 3
	mu.Lock()                // Lock() не может быть вызван дважды, поэтому горутина 3 остановит своё выполнение пока не будет произведён Unlock()
	go Goroutine3(mu, wgAll) // Запуск горутины 3
	time.Sleep(time.Second)
	mu.Unlock() // Продолжение работы горутины 3
	// Горутина 4
	wg := new(sync.WaitGroup) // Используем WaitGroup для приостановки выполнения горутины, она будет ждать, пока значение WaitGroup не будет равно 0
	wg.Add(1)                 // Добавляем значение
	go Goroutine4(wg, wgAll)  // Запуск горутины 4
	time.Sleep(time.Second)
	wg.Done() // Уменьшение значения WaitGroup, горутина 3 продолжает работу
	// Горутина 5
	channelData := make(chan int)     // Канал для передачи данных и остановки выполнения горутины 5
	go Goroutine5(channelData, wgAll) // Запуск горутины 5
	for i := 0; i < 5; i++ {          // Получение данных из горутины
		fmt.Println(<-channelData)
	}
	time.Sleep(time.Second)
	close(channelData) // Закрытие канала выполнит второй блок case в select, произойдёт завершение горутины 4
	wgAll.Wait()       // Ожидание завершения всех горутин
}

func Goroutine1(channel chan int, wgAll *sync.WaitGroup) {
	// Какая-то работа
	fmt.Println("Горутина 1 запущена")
	<-channel // Остановка выполнения горутины пока в канал не поступит значение, либо он не будет закрыт
	// Конец работы горутины
	fmt.Println("Горутина 1 завершает работу")
	wgAll.Done()
}

func Goroutine2(channel chan int, wgAll *sync.WaitGroup) {
	fmt.Println("Горутина 2 запущена")
	channel <- 0 // Запись значения в канал и ожидание чтения из него
	fmt.Println("Горутина 2 завершает работу")
	wgAll.Done()
}

func Goroutine3(mu *sync.Mutex, wgAll *sync.WaitGroup) {
	fmt.Println("Горутина 3 запущена")
	mu.Lock() // Lock() мьютекса, горутина ожидает пока он не станет Unlock()
	fmt.Println("Горутина 3 завершает работу")
	wgAll.Done()
}

func Goroutine4(wg *sync.WaitGroup, wgAll *sync.WaitGroup) {
	fmt.Println("Горутина 4 запущена")
	wg.Wait() // Горутина ожидает пока значение WaitGroup не станет равно 0
	fmt.Println("Горутина 4 завершает работу")
	wgAll.Done()
}

func Goroutine5(channel chan int, wgAll *sync.WaitGroup) {
	fmt.Println("Горутина 5 запущена")
	var num int = 0
	defer fmt.Println("Горутина 5 завершает работу") // Отложенный вызов вывода оповещения о завершении горутины после окончания её работы
	defer wgAll.Done()
	for {
		select { // select для управления работой горутины
		case channel <- num: // если канал открыт, то выполняется данный блок и данные отправляются в канал
			num++
		case <-channel: // при закрытии канала отправка данных невозможна и выполнется этот блок, завершающий работу горутины
			return
		}
	}
} // Также возможна реализация с дополнительным каналом при использовании чтения из канала внутри горутины, при передачи данных в канал для закрытия произойдёт остановка выполнения горутины
