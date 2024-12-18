package main

import "fmt"

func main() {
	var firstName string = "Alif"
	var lastName string

	lastName = "Syah"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)
}

func declareVariableWithoutDataType() string {
	var firstName = "John"

	lastName := firstName

	return lastName
}

func multipleVariable() string {
	var first, second, third string

	first, second, third = "satu", "dua", "tiga"

	one, isFriday, twoPointTwo, say := first, second, third, "Hello"

	allInOne := fmt.Sprintf("%s %s %s %s", one, isFriday, twoPointTwo, say)

	return allInOne
}

func underscoreVariable() string {
	_ = "Belajar Golang"
	_ = "Golang itu gampang"

	name, _ := "Alif", "Ntah Kosong"

	return name
}

func newVariableKeyword() string {
	name := new(string)

	/*
		fmt.Println(name) // will print the pointer value in hexadecimal format
		fmt.Println(*name) // will print the original value from the variable pointer
	*/

	return *name
}

func numericNonDecimalDataType() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
}

func numericDecimalDataType() {
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
}

func boolDataType() {
	var exist bool = true

	fmt.Printf("exist ? %t \n", exist)
}

func constDataType() {
	const firstName string = "Alif"
	const lastName = "Syah"

	fmt.Print("Halo ", firstName, "!\n")
	fmt.Print("Halo ", lastName, "!\n")
}

func constMultipleDataType() string {
	const (
		square   string = "Kotak"
		isToday  bool   = true
		numeric  uint8  = 1
		floatNum        = 2.2
	)

	return square
}

func constSameValueDataType() string {
	const (
		a = "Konstanta"
		b
	)

	const (
		today string = "Senin"
		sekarang
		isToday2 = true
	)

	return a
}

func constMultipleInOneLineDataType() string {
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	return three
}

func conditionIfElse() string {
	var point int = 8

	if point == 10 {
		fmt.Printf("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Printf("lulus")
	} else if point == 4 {
		fmt.Printf("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}

	return "condition if else"
}

func variableTemporaryInCondition() {
	var point = 8840.0

	if percent := point / 100; point >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}

func conditionSwitchCase() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}

func conditionSwitchCaseIfElseStyle() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

}
