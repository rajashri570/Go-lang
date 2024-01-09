// write a program to create the zero value slice slice in go

package main

import "fmt"

func main() {
	var slice []string

	//checing for nil slice
	if slice == nil {
		fmt.Println("its nil slice", slice)
	}

	// assign the array reference to slice

	slice = make([]string, 5)
	slice[0] = "raj"
	slice[1] = "axy"
	slice[2] = "sgf"
	slice[3] = "rai"
	slice[4] = "rsd"

	fmt.Println("slice : ", slice)
	if slice == nil {
		fmt.Println("Its nil slice")
	} else {
		fmt.Println("not nil slice")
	}
}
