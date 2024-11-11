package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64
	//2.вывод типа переменной
	fmt.Printf("Decimal: %d  %T\n", numDecimal, numDecimal)
	fmt.Printf("Octal: %o %T\n", numOctal, numOctal)
	fmt.Printf("Hexadecimal: %x %T\n", numHexadecimal, numHexadecimal)
	fmt.Printf("Pi: %f %T\n", pi, pi)
	fmt.Printf("Name: %s %T\n", name, name)
	fmt.Printf("Active: %t %T\n", isActive, isActive)
	fmt.Printf("Complex: %f %T\n", complexNum, complexNum)
	//преобразовать все в строку и сложить
	res := ""
	a := strconv.Itoa(numDecimal)
	res += a
	a = strconv.Itoa(numOctal)
	res += a
	a = strconv.Itoa(numHexadecimal)
	res += a
	a = strconv.FormatFloat(pi, 'f', -1, 64)
	res += a
	res += name
	a = strconv.FormatBool(isActive)
	res += a
	a = fmt.Sprintf("%v", complexNum)
	res += a
	// 4.сделать из строки руну[]rune(res)
	//5.захэшировать и в середину соль
	str := ""
	str = res[0:len(res)/2] + "go-2024" + res[len(res)/2:]
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
