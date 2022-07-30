package main

import (
	"fmt"
	"sync"
)

type Increment struct {
	i          int
	sync.Mutex // Анонимное встраиваемое поле для использования методов Lock() Unlock() без обращения к полю
}

func (in *Increment) add() { // Метод добавляющие +1 к значению i
	in.Lock() // Lock() мьютекса может быть вызван толкьо один раз, остальные горутины ждут, пока он не будет Unlock()
	in.i++
	in.Unlock()
}

func Work18() {
	wg := new(sync.WaitGroup) // WaitGroup для ожидания завершения всех горутин
	wg.Add(1000)
	obj := new(Increment)
	for i := 0; i < 1000; i++ { // Запускаем 1000 горутин, вызывающих метод add
		go IncrementWorker(obj, wg)
	}
	wg.Wait() // Ждем завершения всех горутин
	fmt.Println(obj.i)
}

func IncrementWorker(in *Increment, wg *sync.WaitGroup) {
	in.add()
	wg.Done() // Завершение работы текущей горутины
}
