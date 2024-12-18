package main

import "fmt"

func main() {
	var fruits = []string{"apple", "manggo", "lemon", "orange"}

	fmt.Println(fruits)

	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	/*
		SLICE IS REFERENCE DATA TYPE, what's mean? if you having on slice and use it in legacy and you try to change one of them the all value will be change (with same reference)
	*/
	baFruits[0] = "pinnaple"

	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	/*
		CAP IS MEAN THE WIDE OF THE SLICE IT SELF,

		EXAMPLE FOR AFRUITS SLICE VALUE TAKE [0:3] IT MEANS SLICE THE VALUE OF FRUITS BECOME [BUAH, BUAH, BUAH, ----] (4 LENGTH),
		BUT WHEN LOOK INTO BFRUITS, THE SLICE VALUE TAKE [1:4] IT MEANS SLICE THE VALUE OF FRUITS BECOME ---- [BUAH, BUAH, BUAH] (3 LENGTH).

		SO FOR THE CONCLUSION THE CAP IS LENGTH OF THE WIDTH OF VALUE SLICE ITSELF, IF THE START FROM 0 INDEX THE CAP LENGTH ALWAYS REFERING TO ORIGINAL LENGTH IF NOT JUST DECREASE WITH THE START INDEX SLICING WITH ORIGINAL LENGTH OF ARRAY
	*/
	// len and cap
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))

	// append
	var oriFruits = []string{"apple", "grape", "banana"}
	var oriBFruits = oriFruits[0:2]

	fmt.Println(cap(oriBFruits))
	fmt.Println(len(oriBFruits))

	fmt.Println(oriFruits)
	fmt.Println(oriBFruits)

	var oriCFruits = append(oriBFruits, "papaya")

	fmt.Println(oriFruits)
	fmt.Println(oriBFruits)
	fmt.Println(oriCFruits)
}
