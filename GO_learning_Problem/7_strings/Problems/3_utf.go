package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	utf_str := "こんにちは, 世界!"
	count := 0

	//itearating through string
	for _, char := range utf_str {

		// increment count if rune value is greater than 127
		if char >= 127 {
			count++
		}
	}

	cnt := 0
	for i := 0; i < len(utf_str); {
		r, size := utf8.DecodeRuneInString(utf_str[i:])
		if r <= 127 {
			fmt.Printf("ASCII character: %c\n", r)
		} else {
			fmt.Printf("UTF-8 character: %c\n", r)
			cnt++
		}
		i += size
	}
	fmt.Println("count : ", count)
	fmt.Println("count without utf: ", cnt)
}
