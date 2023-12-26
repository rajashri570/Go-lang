//write a program to perform to find whether the number is palimdrome or no.
/*algorithm
1.take one numeber
2.store it in one temp variable
3. and then reverse that number
*/
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var num string

	sum := 0
	var err error

	fmt.Println("enter the number : ")

	_, err = fmt.Scan(&num)

	if err != nil {
		fmt.Println("encountered the error: ", err)
	}

	// here i  am checking for the number is having character
	for _, char := range num {
		if !unicode.IsDigit(char) {
			fmt.Println("Error: Please enter a valid number.")
			return
		}
	}

	input, err := strconv.Atoi(num)
	temp := input

	for input > 0 {
		println(input)

		rem := input % 10  // 121 % 10 -> 1
		sum = sum*10 + rem // sum = 0*10 + 10 ->1
		input = input / 10 // num ->1

	}
	if temp == sum {
		fmt.Print(temp, "number is palindrome ")
	} else {
		fmt.Println(temp, "Number is not palindrom ")
	}
}
