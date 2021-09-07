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

func CountDigits(x int) int {
	a := 0
	d := x
	for d > 0 {
		d = d / 10
		a++
	}
	return a
}

func main() {
	var number int

	fmt.Println("Enter an integer number:")
	fmt.Scanf("%d", &number)

	//для любого числа:

	pow := CountDigits(number)
	strFormat := strconv.Itoa(pow)
	result := 0

	for pow > 0 {
		result = result*10 + number%10
		number = number / 10
		pow--
	}

	fmt.Printf("The result is %0"+strFormat+"d\n", result)
}
