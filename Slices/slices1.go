package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(myslice1)) //return length
	fmt.Println(cap(myslice1)) //return capacity
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	//creating a slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	//accessing the element of slice
	fmt.Println(arr1[0])

	//changing the element
	arr1[2] = 67
	fmt.Println(arr1[2])

	//append element in array
	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

}
