package main

import "fmt"

func fibonacci(n int) {
	a := 0
	b := 1
	fmt.Print(a, b, " ")
	for i := 2; i < n; i++ {
		sum := a + b
		fmt.Print(sum, " ")
		a = b
		b = sum
	}

}

func main() {
	var n int
	fmt.Println("Enter the value of n : ")
	fmt.Scan(&n)
	fibonacci(n)
}
