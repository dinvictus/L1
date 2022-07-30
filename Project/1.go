package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Energy int
}

func (obj *Human) DoSomething() { // Метод структуры Human
	obj.Energy--
}

type Action struct {
	NameAction string
	Human      // Анонимное встраиваемое поле структуры Human, позволяет использовать методы Human без обращения к полю
}

func Work1() {
	objAction := Action{NameAction: "Walk", Human: Human{Name: "Dmitry", Age: 21, Energy: 10}}
	fmt.Println("Energy count: ", objAction.Energy)
	objAction.DoSomething() // Вызов метода структуры Human через объект структуры Action
	fmt.Println("Energy count: ", objAction.Energy)
}
