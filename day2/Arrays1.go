package main

import (
	"fmt"
)

func main() {
	//length is defined
	var arr1 = [7]int{1, 2, 3}    //partially initialized
	arr2 := [5]int{4, 5, 6, 7, 8} //fully initialized
	arr5 := [5]int{}              //not initialized

	//inferred/undefined length
	arr3 := [...]int{4, 5, 6, 7, 8}
	var arr4 = [3]int{1, 2, 3}

	//Initialize Only Specific Elements
	arr6 := [5]int{1: 10, 2: 40}

	//strings in array
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(cars)
	fmt.Println("value at second index of arr1:", arr1[2])

	//change value in a index in array
	arr1[2] = 50
	fmt.Println("changed value at second index of arr1:", arr1[2])

	//length of array
	fmt.Println(len(arr2))
}
