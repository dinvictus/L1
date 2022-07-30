package main

import (
	"fmt"
	"sync"
)

type MyMap struct { // Структура для хранения map и мьютекса
	Map        map[int]int
	sync.Mutex // Анонимное поле, чтобы использовать методы мьютекса без обращения к полю
}

func (m *MyMap) Add(key int, value int) { // Метод добавления новых значений в map
	m.Lock()           // Блокировка мьютекса, все остальные горутины ждут, пока он не будет Unlock()
	m.Map[key] = value // Запись новго значения в map
	m.Unlock()         // Другие горутины продолжают работу
}

func Work7() {
	myMap := MyMap{Map: make(map[int]int)} // Создание объекта типа MyMap
	wg := new(sync.WaitGroup)              // Для ожидания завершения всех горутин
	wg.Add(5)                              // Будет запущено 5 горутин
	var start int = 0                      // Число для старта цикла генерации данных для записи
	for i := 0; i < 5; i++ {               // Запуск пяти горутин
		go Worker7(start, &myMap, wg) // Отправка стартового числа, указателя на объект MyMap и WaitGroup
		start += 5
	}
	wg.Wait()              // Ожидание завершения всех горутин
	fmt.Println(myMap.Map) // Вывод map
}

func Worker7(start int, testMap *MyMap, wg *sync.WaitGroup) {
	for i := start; i < start+5; i++ {
		testMap.Add(i, i*i) // Добавление данных в map
	}
	wg.Done() // Горутина завершает свою работу
}
