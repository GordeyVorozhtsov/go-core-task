package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) 
	}
	fmt.Println("original:", arr)
	fmt.Println(sliceExample(arr))     //1.для выборки четных
	fmt.Println(addElements(arr, 5))   //для 2. добавления элемента
	fmt.Println(copySlice(arr))        //для 3. копирования массива
	fmt.Println("original:", arr)      //проверка что массив не изменился в мейне а ток в сopySlice
	fmt.Println(removeElement(arr, 2)) //4.для удаления элемента

}
func sliceExample(arr [10]int) []int {
	newArr := []int{}
	for _, elem := range arr {
		if elem%2 == 0 {
			newArr = append(newArr, elem)
		}
	}
	return newArr
}
func addElements(arr [10]int, elem int) []int {
	newArr := arr[:]
	newArr = append(newArr, elem)
	return newArr
}
func copySlice(arr [10]int) [10]int {
	newArr := [10]int{}
	newArr = arr
	newArr[0] = 666
	return newArr
}
func removeElement(arr [10]int, idx int) []int {
	newArr := arr[:]
	newArr = append(newArr[:idx], newArr[idx+1:]...)
	return newArr
}
