// write a program to create slice from array and iterate with range of slice.

package main

import "fmt"

func main() {

	// created empty array
	var arr = [5]int{}

	arr = [5]int{3, 2, 4, 5, 6}

	fmt.Println("array element : ", arr)
	// iterate the part of array with slice

	// created slice from array ind 1-3r oc
	for _, num := range arr[1:4] {
		fmt.Print(num, " ")
	}
	fmt.Println()

	// this declar slice like arry but without size
	var slice []string
	slice = []string{"rajashri", "Shrikant", "Krishna"}

	// here we have not mentioned start an end as bydefault start = 0 and end len(slice)-1
	fmt.Println(slice[:])

}
