/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var input string
	var number int
	var output string

	fmt.Println("Enter a number of bottles:")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	for number < 0 || number > 200 || err != nil {
		fmt.Println("Not a valid number. Retry, please:")
		fmt.Scanln(&input)
		number, err = strconv.Atoi(input)
	}

	output = "бутылок"

	r := number % 10
	if r > 1 && r < 5 {
		output = strings.ReplaceAll(output, "ок", "ки")
	}
	if r == 1 {
		output = strings.ReplaceAll(output, "ок", "ка")
	}

	fmt.Printf("%d %s\n", number, output)

}
