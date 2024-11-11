package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(Compare(slice1, slice2))

}

func Compare(slice1, slice2 []string) []string {
	largest := []string{}
	shortest := []string{}
	if len(slice1) > len(slice2) {
		largest = slice1
		shortest = slice2
	} else {
		largest = slice2
		shortest = slice1
	}
	compare := make(map[string]int)
	res := []string{}
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
	return res
}
