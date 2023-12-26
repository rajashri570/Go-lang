package main

import (
	"fmt"
)

func main() {
	var input_year int

	fmt.Print("Enter the year : ")
	fmt.Scan(&input_year)

	if input_year%4 == 0 || input_year%100 == 0 {
		fmt.Println("Year is the leap year.")
	} else {
		fmt.Println("Year is not a leap year.")
	}

}
