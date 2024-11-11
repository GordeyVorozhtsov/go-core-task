package main

import (
	"fmt"
)

var StringIntMap = make(map[string]int)

func main() {
	Add("a", 1)
	Add("b", 2)
	Remove("a")
	fmt.Println(StringIntMap)
	fmt.Println(Copy(StringIntMap))
	fmt.Println(Exists("b"))
	fmt.Println(Get("b"))
}
func Add(key string, value int) {
	StringIntMap[key] = value
}
func Remove(key string) {
	delete(StringIntMap, key)
}
func Copy(mapa map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range mapa {
		newMap[k] = v
	}
	return newMap
}
func Exists(key string) bool {
	if _, ok := StringIntMap[key]; ok {
		return true
	}
	return false
}
func Get(key string) (int, bool) {

	if val, ok := StringIntMap[key]; ok {
		return val, true
	}
	return 0, false
}
