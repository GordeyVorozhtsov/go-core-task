package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Compare(a, b))
}

func Compare(slice1, slice2 []int) (bool, []int) {
	largest := []int{}
	shortest := []int{}
	if len(slice1) > len(slice2) {
		largest = slice1
		shortest = slice2
	} else {
		largest = slice2
		shortest = slice1
	}
	compare := make(map[int]int)
	res := []int{}
	for _, v := range largest {
		compare[v] = 0
	}
	for _, v := range shortest {
		if _, ok := compare[v]; ok {
			compare[v]++
		}
	}
	//fmt.Println(compare)
	for idx, v := range compare {
		if v > 0 {
			res = append(res, idx)
		}
	}
	return len(res) > 0, res
}
