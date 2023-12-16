// write a program to find the number which is divisible by 7 but not by 5 in perticular range
//result should be comma seperated.
package main

import (
	"fmt"
)

func main() {
	var slice []int
	for i := 100; i < 200; i++ {
		if i%7 == 0 && i%5 != 0 {
			slice = append(slice, i)
		}
	}
	fmt.Print(slice)
}
