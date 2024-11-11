package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGenerator(ch chan<- int) {
	for {

		num := rand.Intn(10000)
		ch <- num
		time.Sleep(time.Second)
	}
}

func main() {

	ch := make(chan int)

	go randomNumberGenerator(ch)

	for num := range ch {
		fmt.Println("Сгенерировано число:", num)
	}
}
