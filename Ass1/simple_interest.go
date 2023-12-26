package main

import "fmt"

func main() {

	var amt, int_rate, si, year float64

	fmt.Println("Enter the amout : ")
	fmt.Scan(&amt)

	fmt.Println("Enter the interest rate :")
	fmt.Scan(&int_rate)

	fmt.Println("Enter the time in no of years :")
	fmt.Scan(&year)

	si = (amt * int_rate * year) / 100

	fmt.Println("Simple interest is : ", si)

}
