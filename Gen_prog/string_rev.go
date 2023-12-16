// write a program to reverse the string

package main

import (
	"fmt"
)

func main() {
	str := "rajesh"

	for i := len(str) - 1; i > 0; i-- {
		fmt.Printf("%c", str[i])
	}
}
