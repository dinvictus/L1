package main

import (
	"fmt"
)

func Work14() {
	var channelInt interface{} = make(chan int)
	var channelFloat64 interface{} = make(chan float64)
	var str interface{} = "Hello, World!"
	var num interface{} = 123456
	var flag interface{} = false
	var numf interface{} = 123.456
	GetType(channelInt)
	GetType(channelFloat64)
	GetType(str)
	GetType(num)
	GetType(flag)
	GetType(numf)
}

func GetType(obj interface{}) {
	switch el := obj.(type) { // Используем switch по типу
	case int:
		fmt.Println("int:", el)
	case float32:
		fmt.Println("float32:", el)
	case float64:
		fmt.Println("float64:", el)
	case bool:
		fmt.Println("bool:", el)
	case string:
		fmt.Println("string:", el)
	case chan int:
		fmt.Println("Channel int:", el)
	case chan float32:
		fmt.Println("Channel float32:", el)
	case chan float64:
		fmt.Println("Channel float64:", el)
	case chan bool:
		fmt.Println("Channel bool:", el)
	case chan string:
		fmt.Println("Channel string:", el)
	}
}
