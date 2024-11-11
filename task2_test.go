package main

import "testing"

func SliceExample(arr [10]int) []int {
	var result []int
	for _, num := range arr {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}
func AddElements(arr [10]int, elem int) []int {
	newArr := arr[:]
	newArr = append(newArr, elem)
	return newArr
}
func CopySlice(arr [10]int) [10]int {
	newArr := [10]int{}
	newArr = arr
	newArr[0] = 666
	return newArr
}
func RemoveElement(arr [10]int, idx int) []int {
	newArr := arr[:]
	newArr = append(newArr[:idx], newArr[idx+1:]...)
	return newArr
}
func TestSliceExample(t *testing.T) {
	testNums := []struct {
		arr      [10]int
		expected []int
	}{
		{
			[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{2, 4, 6, 8, 10},
		},
		{
			[10]int{62, 93, 37, 99, 47, 55, 62, 7, 28, 58},
			[]int{62, 62, 28, 58},
		},
	}
	for _, testCase := range testNums {
		result := SliceExample(testCase.arr)
		for i, v := range result {
			if v != testCase.expected[i] {
				t.Errorf("For %v, expected %v but got %v", testCase.arr, testCase.expected, result)
				break
			}
		}
		t.Logf("Result: %v", result)
	}
}
func TestAddElements(t *testing.T) {
	testCases := []struct {
		arr      [10]int
		elem     int
		expected []int
	}{
		{
			[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			5,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5},
		},
		{
			[10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			666,
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 666},
		},
	}
	for _, test := range testCases {
		result := AddElements(test.arr, test.elem)
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("For %v, expected %v but got %v", test.arr, test.expected, result)
			}
		}
		t.Logf("Result: %v", result)
	}
}
func TestCopySlice(t *testing.T) { //здесь собсна проверка переданный массив меняется в фукнции а как проверить чтоы он не менялся в мейне хз
	testCases := []struct {
		arr      [10]int
		expected [10]int
	}{
		{
			[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[10]int{666, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[10]int{44, 64, 24, 74, 23, 44, 12, 64, 35, 99},
			[10]int{666, 64, 24, 74, 23, 44, 12, 64, 35, 99},
		},
	}
	for _, test := range testCases {
		result := CopySlice(test.arr)
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("For %v, expected %v but got %v", test.arr, test.arr, result)
			}
		}
		t.Logf("Result: %v", result)
	}
}
