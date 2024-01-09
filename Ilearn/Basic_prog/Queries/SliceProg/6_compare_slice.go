//write a program to compare two strings

package main

import "fmt"

func main() {

	slice1 := []string{"raj", "kishu", "shree"}
	slice2 := []string{"raj", "kishu", "shree"}
	flag := 0

	if len(slice1) != len(slice2) {
		fmt.Println("slice sare not equal")
		return
	}
	for i, val := range slice1 {
		if val != slice2[i] {
			fmt.Println(val, " ", slice2[i])
			fmt.Println("slice are not equal")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("equal slices")
	}

}
