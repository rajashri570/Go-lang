//write a program to calculate the HCF of two numbers

package main

import "fmt"

func get_LCM(num1, num2 int) {

	//var LCM = 1
	var i int
	for i = 1; i <= num1*num2; i++ {
		if i%num1 == 0 && i%num2 == 0 {
			break
		}
	}
	fmt.Printf("LCM of %d and %d is: %d\n", num1, num2, i)
}

func get_HCF(num1, num2 int) {
	var smaller int
	var HCF int
	if num1 < num2 {
		smaller = num1
	} else {
		smaller = num2
	}
	for i := 1; i <= smaller; i++ {
		if num1%i == 0 && num2%i == 0 {
			HCF = i
		}
	}
	fmt.Printf("HCF of %d and %d is: %d\n", num1, num2, HCF)
}
func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers : ")
	fmt.Scan(&num1, &num2)
	get_HCF(num1, num2)
	get_LCM(num1, num2)
}
