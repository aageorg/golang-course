/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
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

	for !((a <= b) && (b <= c)) {
		//       fmt.Println("iteration...")
		if a > b {
			t = a
			a = b
			b = t
		}
		if b > c {
			t = b
			b = c
			c = t
		}

	}
	fmt.Println("The sorted sequence: %d, %d, %d\n", a, b, c)
}
