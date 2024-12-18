package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i+1)
	// }

	// var i = 0

	// while looping (for looping if format)
	// for i < 5 {
	// 	fmt.Println("Angka", i+1, "(while looping)")

	// 	i++
	// }

	// condition looping
	// for {
	// 	fmt.Println("Angka", i+1, "(condition looping)")

	// 	i++

	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for range looping
	var xs = "123"
	for i, v := range xs {
		fmt.Println("Index = ", i, "Value = ", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50}
	for _, v := range ys {
		fmt.Println("Value = ", v)
	}

	var zs = ys[0:2]
	fmt.Println("\n")
	for _, v := range zs {
		fmt.Println("Value = ", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2}
	fmt.Println("\n")
	for k, v := range kvs {
		fmt.Println("Key = ", k, "Value = ", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Println(i) // 01234
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka = ", i)
	}
}
