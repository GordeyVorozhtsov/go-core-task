package main

import (
	"fmt"
	"time"
)

// если семафора это ограничения на количество одновременных горутин,
// то я нашел вот это https://gobyexample.com.ru/rate-limiting
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	for v := range requests {
		fmt.Println("request", v, time.Now())
	}
	moreThanLimit(11)
}

// эта функция тип если задач больше емкости канала и задачи в ожидании
func moreThanLimit(n int) {
	ch := make(chan int, 5)
	go func() {
		for v := range ch {
			fmt.Println("ch", v, time.Now())
		}
	}()
	for l := 0; l < n; {
		for i := 1; i <= 5 && l < n; i++ {
			ch <- i
			l++
		}
	}
	close(ch)
	time.Sleep(1 * time.Second)
}
