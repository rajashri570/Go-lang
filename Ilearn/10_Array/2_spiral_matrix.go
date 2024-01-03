package main

import (
	"fmt"
)

// arr[0]

func main() {

	//  created 3x3 2D matrix
	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// m := len(matrix)    //no of rows
	// n := len(matrix[0]) // no of col

	//declaration of variable to track move
	colBegin := 0
	colEnd := 2
	rowBegin := 0
	rowEnd := 2

	// break condition for forloop
	for rowBegin <= rowEnd && colBegin <= colEnd {

		// print first row with moving left to right
		for j := colBegin; j <= colEnd; j++ {

			fmt.Print(matrix[rowBegin][j], " ")
		}
		rowBegin++ //this is for move down,rowbegin will be 1

		for i := rowBegin; i <= rowEnd; i++ {
			fmt.Print(matrix[i][colEnd], " ")
		}
		colEnd-- // this is for move to left

		if rowBegin <= rowEnd {
			for k := colEnd; k >= colBegin; k-- {
				fmt.Print(matrix[rowEnd][k], " ") //prints the element in reverse maner
			}
		}
		rowEnd-- //this will move to up
		if colBegin <= colEnd {
			for j := rowEnd; j >= rowBegin; j-- {
				fmt.Print(matrix[j][colBegin], " ")
			}
			colBegin++
		}
	}

}
