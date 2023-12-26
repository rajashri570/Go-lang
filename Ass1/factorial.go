package main

import "fmt"

func fact(n int) int {
	if n < 0 {
		fmt.Println("we cant find the factoril of -ve number : ")
		return 0
	} else if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	var num int
	fmt.Println("Enter the value of num : ")
	fmt.Scan(&num)
	res := fact(num)
	fmt.Println("Factorial of ", num, ":", res)
}
