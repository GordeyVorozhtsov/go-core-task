package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	combined := inOnce(ch1, ch2, ch3)
	for val := range combined {
		fmt.Println(val)
	}
}

func inOnce(ch1, ch2, ch3 chan int) <-chan int {
	one := make(chan int)
	go func() {
		defer close(one)
		for val := range ch1 {
			one <- val
		}
		for val := range ch2 {
			one <- val
		}
		for val := range ch3 {
			one <- val
		}
	}()

	return one
}
