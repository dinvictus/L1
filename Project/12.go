package main

import "fmt"

func Work12() {
	Work12_1()
}

func Work12_1() { // 9240 ns/op		5664 B/op		2 allocs/op
	var arrStr []string = []string{"cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree"}
	MapArr := make(map[string]struct{}) // map ?????? ???????????????? ?? ???????????? ???????????????????? ???????????????? ??????????
	for _, el := range arrStr {         // ???????????????? ???? ??????????????
		if _, ok := MapArr[el]; !ok { // ???????? ???????????? ?????????? ???? ??????????????, ???? ?????????????????? ?????????????? ?? ???????? ?? ?? map, ?????????? ?????????????? ?????????????????? ???????????????????? ?? ?????????? ???????????? ???????????????????? ??????????????????
			MapArr[el] = struct{}{}
		}
	}
	fmt.Print(MapArr)
}

func Work12_2() { // 12405 ns/op		5376 B/op		1 allocs/op
	var arrStr []string = []string{"cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree", "cat", "cat", "dog", "cat", "tree", "cat", "dog", "bird", "tree", "pig", "bird", "apple", "cat", "orange", "bird", "dog", "lemon", "pig", "flower", "tree"}
	arrRes := make([]string, 0, len(arrStr))
	for _, el := range arrStr {
		flag := true
		for _, el2 := range arrRes {
			if el == el2 {
				flag = false
			}
		}
		if flag {
			arrRes = append(arrRes, el)
		}
	}
	fmt.Println(arrRes)
}
