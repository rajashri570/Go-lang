package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the value of number : ")
	fmt.Scan(&num)
	fmt.Println("---Multiplication table is ----")
	for i := 1; i <= 10; i++ {

		fmt.Println(i, "*", num, ": ", num*i)

	}
}
