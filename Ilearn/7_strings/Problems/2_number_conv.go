/*
Write a program that takes numeric input as a string,
converts it to a numeric type using the strconv package,
performs a mathematical operation, and then converts the result back to a string for display.
*/

package main

import (
	"fmt"
	"strconv"
)

func increment_number(num int) string {
	num = num + 10
	//afetr increment converting integer to string again
	return strconv.Itoa(num)
}
func main() {
	var input string
	fmt.Println("Enter the number :")
	fmt.Scan(&input)

	//converting input of string to int
	number, err := strconv.Atoi(input)

	// handling error in case if input having letter or symbols
	if err != nil {
		fmt.Println("Invalid input!\nPlease Enter valid input")
		return
	} else {
		fmt.Printf("\nNumber %v is of type %T\n\n", number, number)
	}

	//increment number by 10
	res := increment_number(number)
	fmt.Printf("Result of increment by 10 :%s\n", res)

}
