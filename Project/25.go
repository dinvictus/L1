package main

import (
	"fmt"
	"time"
)

func Work25() {
	start := time.Now()
	MySleep_1(500 * time.Millisecond)
	fmt.Println(time.Since(start))
}

func MySleep_1(d time.Duration) {
	<-time.After(d) // Горутина спит, пока в канал не поступит значение через заданное время
}
