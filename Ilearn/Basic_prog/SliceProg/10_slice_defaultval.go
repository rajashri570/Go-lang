//write a program to create a slice with new opearator

package main

import (
	"fmt"
	"reflect"
)

func main() {

	// create a slice variable

	var slice = new([20]int)
	var arr [5]int = [5]int{2, 34, 5, 67, 8}
	fmt.Println("Length :", len(slice))
	fmt.Println("Capacity :", cap(slice))
	fmt.Println(reflect.ValueOf(arr).Kind())
	//fmt.Printf("%T", slice)

	fmt.Println(slice, arr)
}

/* rdolzake@rdolzake-HP-245-G5-Notebook-PC:~/GOLANG/Go-lang/Ilearn/Basic_prog/Queries/SliceProg$ go run 10_slice_defaultval.go
Length : 10
Capacity : 20
[0 0 0 0 0 0 0 0 0 0]
*/
