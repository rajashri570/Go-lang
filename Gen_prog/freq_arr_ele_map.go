//write a program to get one array and print the count of elements occured in that array.

package main

import "fmt"

func main() {
	var arr = [7]int{3, 1, 2, 5, 3, 2, 5}
	var arr_map = make(map[int]int)

	for _, v := range arr {
		// this line checks whether the key is present in arr_map
		//if found = true then just increament count value else add 1 as count
		if count, found := arr_map[v]; found {

			arr_map[v] = count + 1

		} else {

			arr_map[v] = 1
		}

	}
	fmt.Println("Array and count :")
	fmt.Print(arr_map)

}
