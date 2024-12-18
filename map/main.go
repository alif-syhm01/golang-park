package main

import "fmt"

func main() {
	// init map
	var chicken map[string]int

	// add bracket for avoiding nil value
	chicken = map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	// fmt.Println("Januari", chicken["Januari"])
	// fmt.Println("Mei", chicken["Mei"])

	var chicken1 = map[string]int{
		"januari":  50,
		"februari": 23,
		"maret":    22,
		"april":    41,
		"mei":      40,
	}

	// fmt.Println("chicken 1", chicken1)

	// another way to make map without curly braces
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	for key, val := range chicken1 {
		fmt.Println(key, " \t:", val)
	}

	// delete the map
	delete(chicken1, "maret")

	fmt.Println(len(chicken1))
	fmt.Println(chicken1)

	// check the item with key in map
	var value, isExist = chicken1["april"]

	fmt.Println(value, isExist)

	var students = []map[string]string{
		{"name": "Andi", "gender": "male"},
		{"name": "Putri", "gender": "female"},
		{"name": "Bayu", "gender": "male"},
	}

	for _, student := range students {
		fmt.Println(student["name"], student["gender"])
	}
}
