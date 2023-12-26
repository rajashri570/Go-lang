package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("Enter the string :")
	fmt.Scan(&str)
	revstr := ""

	for i := len(str) - 1; i >= 0; i-- {
		revstr = revstr + string(str[i])
	}
	fmt.Println("reversed string: ", revstr)

	//With other method
	//converting string to the rune array to locate the each char
	rune := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {

		rune[i], rune[j] = rune[j], rune[i]

	}
	fmt.Println("Reversed string: ", string(rune))
}
