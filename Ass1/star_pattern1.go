/*
write a program to print the output like :

*           // 0
* *			//1
* * *		//2
* * * *		//3
* * * * *	// 4
* * * *		//5
* * *		//6
* *			//7
*			//8

Algorithm :
1.first we will decide the no of time outer loop will run i.e 9 times So row = 0; rao < 9 ;i ++
2. decide the mid till what you hav to consider colum  = row
3. after reaching to mid col = row - mid
*/

package main

import "fmt"

func main() {
	n := 9
	mid := n / 2                   //  4
	for row := 0; row < 9; row++ { // 0 1 2 3 4 5
		if row <= mid {
			for col := 0; col <= row; col++ {
				fmt.Print("*", " ")
			}
			fmt.Println()
		} else {
			for col := 0; col < n-row; col++ { // 9 - 5 ->4
				fmt.Print("*", " ")
			}
			fmt.Println()
		}
	}
}
