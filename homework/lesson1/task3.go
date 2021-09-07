/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var c int

	var t int

	fmt.Println("Enter the first integer number:")
	fmt.Scanf("%d", &a)

	fmt.Println("Enter the second integer number:")
	fmt.Scanf("%d", &b)

	fmt.Println("Enter the the third integer number:")
	fmt.Scanf("%d", &c)

	for a > b || b > c || c <= a {
		//       fmt.Println("iteration...")
		if a >= b {
			t = a
			a = b
			b = t
		}
		if b >= c {
			t = b
			b = c
			c = t
		}
		if c <= a {
			t = c
			c = a
			c = t
		}
	}
	fmt.Println("The sorted sequence: %d, %d, %d\n", a, b, c)
}
