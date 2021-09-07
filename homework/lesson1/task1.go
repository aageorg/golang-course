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

	const price = 48

	var distance int
	var consumption float32
	var total float32

	fmt.Println("Enter the distance length, km (50-50000):")
	fmt.Scanf("%d", &distance)

	fmt.Println("Enter the estimated fuel consumption, l/100km:")
	fmt.Scanf("%f", &consumption)

	if distance >= 50 && distance <= 50000 {

		total = consumption / 100 * float32(distance) * price

		fmt.Printf("The total price is %.2f RUB\n", total)

	} else {
		fmt.Println("Unsupported distance value\n")
	}

}
