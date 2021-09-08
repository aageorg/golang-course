/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
	"fmt"
	"math"
	//	"strconv"
	//	"strings"
)

func TriangleAble(xy [3][2]float64) bool {

	if math.Round((xy[2][0]-xy[0][0])/(xy[1][0]-xy[0][0])) == math.Round((xy[2][1]-xy[0][1])/(xy[1][1]-xy[0][1])) {
		return false
	} else {
		return true
	}
}

func TriangleSides(xy [3][2]float64) (float64, float64, float64) {
	a := math.Sqrt(math.Pow((xy[1][0]-xy[0][0]), 2) + math.Pow((xy[1][1]-xy[0][1]), 2))
	b := math.Sqrt(math.Pow((xy[2][0]-xy[1][0]), 2) + math.Pow((xy[2][1]-xy[1][1]), 2))
	c := math.Sqrt(math.Pow((xy[2][0]-xy[0][0]), 2) + math.Pow((xy[2][1]-xy[0][1]), 2))
	return a, b, c
}

func STriangle(a float64, b float64, c float64) float64 {

	p := (a + b + c) / 2

	return math.Sqrt(p * (p - a) * (p - b) * (p - c))

}

func IsRectangular(a float64, b float64, c float64) bool {

	var maxSide float64
	var cat1 float64
	var cat2 float64

	if a > b && a > c {
		maxSide = a
		cat1 = b
		cat2 = c
	} else if b > a && b > c {
		maxSide = b
		cat1 = a
		cat2 = c
	} else if c > a && c > b {
		maxSide = c
		cat1 = a
		cat2 = b
	} else {
		return false
	}
	if math.Round(math.Pow(maxSide, 2)) == math.Round((math.Pow(cat1, 2) + math.Pow(cat2, 2))) {
		return true
	} else {
		return false
	}
}

func main() {

	var x float64
	var y float64
	var rect string
	var coordinates [3][2]float64

	fmt.Println("Enter the first x and y coordinates (e.g. 2 3):")
	count, _ := fmt.Scanf("%f %f", &x, &y)
	for count != 2 || x < 0 || y < 0 {
		fmt.Println("Incorrect input, retry, please:")
		count, _ = fmt.Scanf("%f %f", &x, &y)
	}
	coordinates[0][0] = x
	coordinates[0][1] = y

	fmt.Println("Enter the second x and y coordinates (e.g. 2 3):")
	count, _ = fmt.Scanf("%f %f", &x, &y)
	for count != 2 || x < 0 || y < 0 {
		fmt.Println("Incorrect input, retry, please:")
		count, _ = fmt.Scanf("%f %f", &x, &y)
	}
	coordinates[1][0] = x
	coordinates[1][1] = y

	fmt.Println("Enter the third x and y coordinates (e.g. 2 3):")
	count, _ = fmt.Scanf("%f %f", &x, &y)
	for count != 2 || x < 0 || y < 0 {
		fmt.Println("Incorrect input, retry, please:")
		count, _ = fmt.Scanf("%f %f", &x, &y)
	}
	coordinates[2][0] = x
	coordinates[2][1] = y

	if TriangleAble(coordinates) == true {

		fmt.Println("По заданным точкам можно построить треугольник")
		a, b, c := TriangleSides(coordinates)
		fmt.Printf("Площадь треугольника будет %f\n", STriangle(a, b, c))
		fmt.Printf("Стороны треугольника будут %f %f %f\n", a, b, c)
		isRectangular := IsRectangular(a, b, c)
		if isRectangular == true {
			rect = "да"
		} else {
			rect = "нет"
		}

		fmt.Printf("Является ли прямоугольным: %s\n", rect)

	} else {

		fmt.Println("По заданным точкам нельзя построить треугольник")

	}
}
