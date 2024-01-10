// write a program to create the slice with different kinds of data  and
// print the type of data stored in slice

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var mySlice []interface{}
	fmt.Println(mySlice...)

	mySlice = append(mySlice, 42, "rajashir", 3.14)

	fmt.Println("slice values are :", mySlice)

	for _, val := range mySlice {

		fmt.Println("value getted with valueOf:", reflect.ValueOf(val))
		fmt.Println(val, "is of ", reflect.ValueOf(val).Kind(), "type")
	}

}
