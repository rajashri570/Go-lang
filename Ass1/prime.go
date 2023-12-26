//write a program to check whether the number is palindrome or no :

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the value : ")
	fmt.Scan(&num)
	var flag = 0
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Print(num, " is not prime number")
	} else {
		fmt.Print(num, " is prime number")
	}
}
