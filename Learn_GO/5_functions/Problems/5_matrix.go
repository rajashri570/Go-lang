/*
Matrix Multiplication:

Implement a function that performs matrix multiplication.
 Given two matrices as input,
 return the resulting matrix.
 Make sure to check the compatibility of the matrices for multiplication.


package main

import "fmt"

func main() {
	var no_row int
	var no_col int

	fmt.Println("Enter the number of rows in matrics 1: ")
	fmt.Scan(&no_row)
	fmt.Println("Enter the number of column in matrics 1: ")
	fmt.Scan(&no_col)

	//fmt.Println("Enter the value for ", no_row, "*", no_col, "Matrix")
	matrix1 := make([][]int, no_row)
	for i := range matrix1 {
		matrix1[i] = make([]int, no_col)
	}

	fmt.Println("Enter the value for ", no_row, "*", no_col, "Matrix")
	for i := 0; i < no_row; i++ {
		for j := 0; j < no_col; j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}

	//declaring the variable to hold no of columns and rows in matrix2
	var no_col2 int
	var no_row2 int

	fmt.Println("Enter the number of rows in matrics 2: ")
	fmt.Scan(&no_row2)
	fmt.Println("Enter the number of column in matrics 2: ")
	fmt.Scan(&no_col2)

	matrix2 := make([][]int, no_row2)
	for i := range matrix2 {
		matrix2[i] = make([]int, no_col2)
	}

	fmt.Println("Enter the value for ", no_row2, "*", no_col2, "Matrix")

	for i := 0; i < no_row2; i++ {
		for j := 0; j < no_col2; j++ {
			fmt.Scan(&matrix2[i][j])
		}
	}

	fmt.Println("-----------Matrix1-------------")
	for i := 0; i < no_row; i++ {
		for j := 0; j < no_col; j++ {
			fmt.Print(matrix1[i][j], "  ")

		}
		fmt.Println()

	}

	fmt.Println("-----------Matrix2-------------")
	for i := 0; i < no_row2; i++ {
		for j := 0; j < no_col2; j++ {
			fmt.Print(matrix2[i][j], "  ")

		}
		fmt.Println()

	}

	if no_col == no_row2 {
		fmt.Println("both matrics compitable for matrix multiplication ")
	} else {
		fmt.Println("Matrics are not compitable for multiplication")
	}

	rows1, cols1 := len(matrix1), len(matrix1[0])
	rows2, cols2 := len(matrix2), len(matrix2[0])

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}
	for i := 0; i < no_row; i++ {
		for j := 0; j < no_col2; j++ {
			for k := 0; k < no_col; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

}
*/
package main

import "fmt"

func main() {
	var no_row1 int
	var no_col1 int

	// Input for matrix 1
	fmt.Println("Enter the number of rows in matrix 1: ")
	fmt.Scan(&no_row1)
	fmt.Println("Enter the number of columns in matrix 1: ")
	fmt.Scan(&no_col1)

	matrix1 := make([][]int, no_row1)
	for i := range matrix1 {
		matrix1[i] = make([]int, no_col1)
	}

	fmt.Println("Enter the values for matrix 1:")
	for i := 0; i < no_row1; i++ {
		for j := 0; j < no_col1; j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}

	// Input for matrix 2
	var no_row2 int
	var no_col2 int

	fmt.Println("Enter the number of rows in matrix 2: ")
	fmt.Scan(&no_row2)
	fmt.Println("Enter the number of columns in matrix 2: ")
	fmt.Scan(&no_col2)

	if no_col1 != no_row2 {
		fmt.Println(" matrices are not compatible for multiplication.")
		return
	}

	matrix2 := make([][]int, no_row2)
	for i := range matrix2 {
		matrix2[i] = make([]int, no_col2)
	}

	fmt.Println("Enter the values for matrix 2:")
	for i := 0; i < no_row2; i++ {
		for j := 0; j < no_col2; j++ {
			fmt.Scan(&matrix2[i][j])
		}
	}

	// Perform matrix multiplication and print result
	fmt.Println("----Resultant Matrix-----")
	for i := 0; i < no_row1; i++ {
		for j := 0; j < no_col2; j++ {
			var result int
			for k := 0; k < no_col1; k++ {
				result += matrix1[i][k] * matrix2[k][j]
			}
			fmt.Print(result, "  ")
		}
		fmt.Println()
	}
}
