package main

import "fmt"

func main() {
	var inputWork string
	for loop := true; loop; {
		fmt.Scan(&inputWork)
		switch inputWork {
		case "1":
			Work1()
		case "2":
			Work2()
		case "3":
			Work3()
		case "4":
			Work4()
		case "5":
			Work5()
		case "6":
			Work6()
		case "7":
			Work7()
		case "8":
			Work8()
		case "9":
			Work9()
		case "10":
			Work10()
		case "11":
			Work11()
		case "12":
			Work12()
		case "13":
			Work13()
		case "14":
			Work14()
		case "15":
			Work15()
		case "16":
			Work16()
		case "17":
			Work17()
		case "18":
			Work18()
		case "19":
			Work19()
		case "20":
			Work20()
		case "21":
			Work21()
		case "22":
			Work22()
		case "23":
			Work23()
		case "24":
			Work24()
		case "25":
			Work25()
		case "26":
			Work26()
		case "q":
			loop = false
		default:
			fmt.Println("Incorrect input")
		}
	}
}
