/*
Write a program to print the pattern like below :
A
A  B
A  B  c
A  B  C  D
algo:

1. It will need two loops one to store rows and other for column
*/

package main

import "fmt"

func main() {
	for i := 'A'; i <= 'D'; i++ {
		for j := 'A'; j <= i; j++ {
			fmt.Printf("%c  ", j)
		}
		fmt.Println()
	}

}
