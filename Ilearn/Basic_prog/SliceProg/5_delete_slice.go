//write a program to delete the element from slice
//ans :there is no builtin function for deleting so we can use subslices

package main

import "fmt"

func main() {

	slice := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Before deletion: ", slice)

	index_remove := 2
	newSlice := append(slice[:index_remove], slice[index_remove+1:]...)
	fmt.Println("Afer deletion slice :", newSlice)
}
