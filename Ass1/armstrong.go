package main

import (
	"fmt"
	"math"
)

func armstrong(num int) bool {
	//this for counting digits in number
	no := num
	cnt := 0
	for no > 0 {
		no = no / 10
		cnt = cnt + 1
	}
	//fmt.Print(cnt)

	temp := num
	sum := 0
	for num > 0 {
		rem := num % 10
		sum = sum + int(math.Pow(float64(rem), float64(cnt)))
		num = num / 10
	}
	if temp == sum {
		return true
	} else {
		return false
	}

}

func main() {
	var num int
	fmt.Println("Enter the value of number : ")
	fmt.Scan(&num)

	if armstrong(num) {
		fmt.Println(num, " is armstrong number.")
	} else {
		fmt.Println("Number is not armstrong number")
	}

}
