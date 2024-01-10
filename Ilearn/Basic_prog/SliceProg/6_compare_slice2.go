package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice1 := []string{"raj", "kishu", "shree"}
	slice2 := []string{"raj", "kishu", "shree"}
	slice3 := []string{"aaa", "sss"}

	if len(slice1) != len(slice2) {
		fmt.Println("slice are not equal")
		return
	}

	if reflect.DeepEqual(slice1, slice2) {
		fmt.Println("slice1 ans slice2 are equal")
	} else {
		fmt.Println("Slice1 and slice2 not equal")
	}

	fmt.Println(reflect.DeepEqual(slice1, slice3))
}
