package main

import "fmt"

func main() {
	// init array with empty value
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	// init array with values
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// init array without length
	var numbers = [...]int{1, 2, 3, 4, 5}

	fmt.Println("Jumlah element \t", len(numbers))
	fmt.Println("Isi semua element \t", numbers)

	// multidimensional array
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{7, 8, 9}, {10, 11, 12}} // better optional

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// loop with arrays
	var fruitsInArray = [4]string{"apple", "manggo", "orange", "lemon"}

	// with len
	for i := 0; i < len(fruitsInArray); i++ {
		fmt.Printf("elemen %d: %s\n", i, fruitsInArray[i])
	}

	// with range
	for i, fruit := range fruitsInArray {
		fmt.Printf("elemen %d: %s (range)\n", i, fruit)
	}

	// with range and unused index
	for _, fruit := range fruitsInArray {
		fmt.Printf("nama buah: %s (range)\n", fruit)
	}

	var makeFruits = make([]string, 2)

	makeFruits[0] = "apple"
	makeFruits[1] = "manggo"

	fmt.Println(makeFruits)
}
