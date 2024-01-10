package main

import "fmt"

func main() {

	//declaration of slice1 and slice2
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6}

	//create slice with addition length of slice1 and slice2
	size := len(slice1) + len(slice2)
	result := make([]int, size)

	//copy first slice to result

	copy(result, slice1)

	//copy second slice to resultconst
	//first reached to end of slice1 in result
	copy(result[len(slice1):], slice2)

	fmt.Println(result)

	resultSlice := result
	fmt.Println("copy with assignmnt:", resultSlice)

}
