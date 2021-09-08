package main

import (
	"fmt"
)

func Dimentions(a float64, b float64) (float64, float64) {

	s := a * b
	p := (a + b) * 2

	return s, p
}

func BorderPrice(p float64, meterPrice float64) float64 {
	return p * meterPrice
}

func LandPrice(s float64, squarePrice float64) float64 {
	return s * squarePrice * 1.3
}

func main() {

	var meterPrice float64
	var squarePrice float64
	var a float64
	var b float64

	fmt.Println("Длина участка:")
	fmt.Scanf("%f", &a)
	fmt.Println("Ширина участка:")
	fmt.Scanf("%f", &b)
	fmt.Println("Стоимость квадратного метра земли:")
	fmt.Scanf("%f", &meterPrice)
	fmt.Println("Стоимость погонного метра забора:")
	fmt.Scanf("%f", &squarePrice)

	s, p := Dimentions(a, b)

	landPrice := LandPrice(s, squarePrice)
	borderPrice := BorderPrice(p, meterPrice)

	fmt.Printf("Стоимость участка земли: %.2f\n", borderPrice)
	fmt.Printf("Стоимость забора по периметру: %.2f\n", landPrice)

}
