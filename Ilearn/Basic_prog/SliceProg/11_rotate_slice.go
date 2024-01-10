// write a program to rotate left the slice by given position.

package main

import "fmt"

func main() {

	org_slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(org_slice)
	newslice := make([]int, 6)
	pos := 2

	//used this to keep track of elements in new slice
	count := 0

	//coping the element form original slice to new slice

	for i := 0; i < len(org_slice)-pos; i++ {

		newslice[i] = org_slice[i+pos]
		count++

	}
	// making subslice of element to shift at the end of new slice
	sub := org_slice[:pos]
	for _, val := range sub {
		fmt.Print(val)
		newslice[count] = val
		count++
	}
	//newslice = append(newslice, sub...)
	fmt.Println("Original slice: ", org_slice)
	fmt.Println("after roration slice :", newslice)

}
