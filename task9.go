package main

import (
	"fmt"
	"time"
)

func main() {
	i := []uint8{1, 2, 3, 4, 5}
	ch := make(chan uint8)
	go func() {
		for v := range ch {

			firstCH(v)
		}
	}()
	for _, v := range i {
		ch <- v
	}
	close(ch)
	time.Sleep(50 * time.Millisecond) //если без этой задеркжи, то куб пяти не выводится,
	//если добавить в массив после пяти - шесть, то выведется и куб пяти и куб шести, решил оставить задержку
}

func firstCH(n uint8) {
	sc := make(chan float64)
	var floatNum float64 = float64(n)
	go func() {
		sc <- floatNum * floatNum * floatNum
	}()
	secondCH(sc)
}
func secondCH(ch chan float64) {
	k := <-ch
	fmt.Printf("type %f is %T\n", k, k)
}
