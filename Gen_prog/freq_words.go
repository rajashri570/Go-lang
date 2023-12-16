//write a program to get line of text and find the frequency of words in that string

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	word_map := make(map[string]int)

	fmt.Println("enter the line of text :")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error occured reading %s", line)
	} else {
		//count := 1
		word_slice := strings.Fields(line)
		for _, v := range word_slice {
			count := 1
			if _, found := word_map[v]; found {

				word_map[v] = count + 1
			} else {
				word_map[v] = count
			}
		}
	}
	fmt.Println()
	for key, _ := range word_map {
		fmt.Println(key, ":", word_map[key])
	}
}
