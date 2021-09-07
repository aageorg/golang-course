/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import (
	"fmt"
)

func main() {

	var numberX int
	var reverseX int
	var temp int

	fmt.Println("Enter a 4-digit integer number:")
	fmt.Scanf("%d", &temp)

	for temp < 999 || temp > 9999 {
		fmt.Println("Number is not a 4-digit integer number. Retry, please:")
		fmt.Scanf("%d", &temp)

	}

	numberX = temp

	reverseX = numberX/1000 + numberX/100%10*10 + numberX/10%10*100 + numberX%10*1000
	if numberX == reverseX {
		fmt.Printf("%d is a palindome \n", numberX)
	} else {
		fmt.Printf("%d is not a palindrome\n", numberX)
	}
}
