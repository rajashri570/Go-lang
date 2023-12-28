package main

import (
	"fmt"
)

func main() {
	// A string containing ASCII and Unicode characters
	str := "Go世界"

	fmt.Println("String:", str)

	// Print Unicode code points
	fmt.Println("Rune Code Points:")
	for _, runeValue := range str {
		fmt.Printf("%c: %U\n", runeValue, runeValue)
	}

	// Print UTF-8 encoded bytes
	fmt.Println("UTF-8 Bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c: %x\n", str[i], str[i])
	}
}
