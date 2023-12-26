/* write a program to print output
		A
	   A B
	  A B C
	A  B  C D





/*
write a program to print the output like :

					*
				*		*
			*		*		*
		*		*		*		*
	*		*		*		*		*
		*		*		*		*
			*		*		*
				*		*
					*


*/

package main

import "fmt"

func main() {
	n := 9
	mid := n / 2                   //  4
	for row := 0; row < n; row++ { // 0 1 2 3 4 5

		if row <= mid {

			spaces := mid - row

			for k := 0; k < spaces; k++ {
				fmt.Print(" ")
			}
			for col := 0; col <= row; col++ {
				fmt.Print("*", " ")
			}
			fmt.Println()
		} else {
			spaces := row - mid

			for k := 0; k < spaces; k++ {
				fmt.Print(" ")
			}
			for col := 0; col < n-row; col++ { // 9 - 5 ->4
				fmt.Print("*", " ")
			}
			fmt.Println()
		}
	}
}
