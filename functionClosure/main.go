package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (min int, max int) {
		for i, currNumber := range n {
			switch {
			case i == 0:
				max, min = currNumber, currNumber
			case currNumber > max:
				max = currNumber
			case currNumber < min:
				min = currNumber
			}
		}

		return
	}

	var numbers = []int{2, 3, 4, 2, 3, 4}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	var muchNumbers = []int{2, 1, 2, 3, 2, 4, 5, 5, 4, 1, 2, 3, 4, 4, 5, 1}

	// Immediately Invoked Function Expression (IIFE)
	var newNumbers = func(min int) []int {
		var r []int

		for _, e := range muchNumbers {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("Original number: ", muchNumbers)
	fmt.Println("Filtered number: ", newNumbers)

	// var closure (func (string, int, []string) int)
	// closure = func (a string, b int, c []string) int {
	//     // ..
	// }

	var maxNumber = 3
	var listNumbers = []int{2, 3, 1, 2, 0, 4, 5, 6, 6, 7, 8, 2, 1, 3, 3}
	var howMany, getNumbers = findMax(listNumbers, maxNumber)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t: ", listNumbers)
	fmt.Println("find \t: ", maxNumber)

	fmt.Println("found \t: ", howMany)
	fmt.Println("value \t: ", theNumbers)
}

// closure with return function
func findMax(numbers []int, max int) (int, func() []int) {
	var maxNumbers []int

	for _, e := range numbers {
		if e <= max {
			maxNumbers = append(maxNumbers, e)
		}
	}

	var getList = func() []int {
		return maxNumbers
	}

	return len(maxNumbers), getList
}
