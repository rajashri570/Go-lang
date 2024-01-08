package main

import (
	"BubbleModule/Bubble"
	"fmt"
	"os"
)

type CompareFunc func(a, b rune) bool

func main() {
	// var name = "rajashri"
	// str := []rune(name)
	var input_str string
	var ch int
	//annonymous function for ascending order
	asc := func(a, b rune) bool {
		return a > b
	}

	//annonymous function for Descending order
	des := func(a, b rune) bool {
		return a < b
	}

	for {
		fmt.Println("-------Menu---------")
		fmt.Println("1.Ascending Order")
		fmt.Println("2.Descending Oredr")
		fmt.Println("3.Quit")
		fmt.Println("Enter your choice :")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			fmt.Println("Enter the string : ")
			fmt.Scan(&input_str)
			str := []rune(input_str)
			arr := Bubble.Bubble_sort(str, asc)
			fmt.Println("sorted desc string:", string(arr))
		case 2:
			fmt.Println("Enter the string : ")
			fmt.Scan(&input_str)
			str := []rune(input_str)
			arr := Bubble.Bubble_sort(str, des)
			fmt.Println("sorted desc string :", string(arr))
		case 3:
			os.Exit(0)

		default:
			fmt.Println("invalid choice")

		}
	}

}
