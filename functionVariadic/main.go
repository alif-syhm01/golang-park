package main

import (
	"fmt"
	"strings"
)

func main() {
	// var avg = calculate(2, 3, 4, 5, 3, 1, 2, 23, 1, 2)
	var numbers = []int{2, 3, 4, 5, 3, 1, 2, 23, 1, 2} // with slice

	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)

	// yourHobbies("Alif", "rebahan", "makan", "maen ML sambil toxic (solo)")

	var hobbies = []string{"rebahan", "makan", "maen ML sambil toxic (solo)"}
	yourHobbies("Alif", hobbies...) // with slice
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func calculate(numbers ...int) float64 {
	fmt.Println(numbers)

	var total int = 0

	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))

	return avg
}
