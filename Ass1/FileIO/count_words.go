// write a program to print the no of duplicate numbers in file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filenm string
	var wordMap = make(map[string]int)

	fmt.Println("Enter the file name with .txt extension:")
	fmt.Scan(&filenm)

	file, err := os.Open(filenm)
	if err != nil {
		fmt.Println("FILE OPENING ERROR:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimRight(line, "\n")
		wordList := strings.Fields(line)

		for _, word := range wordList {
			wordMap[word]++
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Println("ERROR READING FILE:", err)
		return
	}

	//fmt.Println("Word frequency:")
	fmt.Println("\ndupicate words :")
	for word, freq := range wordMap {
		//fmt.Printf("%s: %d\n", word, freq)

		if freq > 1 {
			fmt.Println(word, ":", freq)
		}
	}

}
