// write a programs to compare two maps in go

package main

import (
	"fmt"
	"reflect"
)

func compare(map1, map2 map[int]string) bool {

	if len(map1) != len(map2) {

		return false

	} else {

		for key, val1 := range map1 {
			val2, ispresent := map2[key]
			if ispresent {
				if val1 == val2 {
					return true
				}
			} else {
				return false
			}

		}

	}
	return false
}

func main() {

	map1 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	map2 := map1
	map3 := map[int]string{1: "monday", 2: "tuesday", 3: "thirsday", 4: "friday", 5: "saturday"}
	fmt.Print(map2)

	if reflect.DeepEqual(map1, map2) {
		fmt.Println("Equal maps")
	} else {
		fmt.Println("Not equal maps")
	}

	fmt.Println("\nWithout built in function :")

	iscompare := compare(map1, map3)
	if iscompare {
		fmt.Println("equals")
	} else {
		fmt.Println("not equal")
	}
}
