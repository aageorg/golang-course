/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import (
	"fmt"
	"strconv"
)

func IntPow(x int, y int) int {
	pow := 1
	a := x
	for pow < y {
		a = a * x
		pow++
	}
	return a
}

func CountDigits(x int) int {
	a := 0
	for x/10 > 0 {
		x = x / 10
		a++
	}
	return a + 1
}

func main() {
	var number int

	fmt.Println("Enter an integer number:")
	fmt.Scanf("%d", &number)

	//для любого числа:

	pow := CountDigits(number) - 1
	sf := strconv.Itoa(pow + 1)
	result := 0

	for CountDigits(number) > 1 {

		lastdigit := number % 10
		result += lastdigit * IntPow(10, pow)
		pow--
		number = number / 10
	}

	result += number

	fmt.Printf("The result is %0"+sf+"d\n", result)
}
