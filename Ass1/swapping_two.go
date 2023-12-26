package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("Enter the value of number1 and number2: ")
	fmt.Scan(&num1, &num2)

	fmt.Println("Before swapping: ")
	fmt.Println("num1 :", num1, "\tnum2 :", num2)

	num1, num2 = num2, num1

	fmt.Println("\nAfter swapping: ")
	fmt.Println("num1 :", num1, "\tnum2 :", num2)

}
