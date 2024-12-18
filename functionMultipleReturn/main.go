package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	fmt.Println("\nPredefined return function")

	var pArea, pCircumreference = calculatePredifendReturn(diameter)

	fmt.Printf("pre luas lingkaran\t\t: %.2f \n", pArea)
	fmt.Printf("pre keliling lingkaran\t\t: %.2f \n", pCircumreference)
}

func calculate(d float64) (float64, float64) {
	// calculate the area
	var area = math.Pi * math.Pow(d/2, 2)

	// calculate the perimeter
	var circumference = math.Pi * d

	return area, circumference
}

func calculatePredifendReturn(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
