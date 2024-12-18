package main

import (
	"fmt"
	"strings"
)

// naming the function scheme
type FilterCallback func(string) bool

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Original Data \t\t\t: ", data)

	fmt.Println("Data Containing Letter \"o\"\t: ", dataContainsO)

	fmt.Println("Data Total Length \"5\"\t\t: ", dataLength5)
}

func filter(data []string, callback FilterCallback) []string {
	var results []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			results = append(results, each)
		}
	}

	return results
}

// func filter(data []string, callback func(string) bool) []string {
// 	var results []string

// 	for _, each := range data {
// 		if filtered := callback(each); filtered {
// 			results = append(results, each)
// 		}
// 	}

// 	return results
// }
