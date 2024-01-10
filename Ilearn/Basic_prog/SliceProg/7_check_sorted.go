package main

import (
	"fmt"
)

func main() {
	var i int
	slice1 := []int{2, 13, 50, 96}
	//flag :=
	for i = 1; i < len(slice1)-1; i++ {
		if !(slice1[i] >= slice1[i-1] && slice1[i] <= slice1[i+1]) {
			fmt.Println("not sorted slice")
			break
		}
	}

	//rev := slice1[::-1]
	//fmt.Println(rev)
	desc_slice := []int{}
	//fmt.Println(desc_slice)
	desc_slice = make([]int, len(slice1))
	for i, v := range slice1 {
		desc_slice[len(slice1)-i-1] = v
		fmt.Println(desc_slice)

	}
	// //fmt.Println(desc_slice)
	// if reflect.DeepEqual(slice1, desc_slice[]) {
	// 	fmt.Println("sorted slice")
	// } else {
	// 	fmt.Println("Not sorted list")
	// }
}
