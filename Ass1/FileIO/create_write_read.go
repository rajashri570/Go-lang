// write a program to create file and write into file and line by line.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file_create.txt")
	if err != nil {
		fmt.Println("Error in file creation:", err)
		return
	}
	defer file.Close()

	fmt.Println(file.Name(), "is created successfully")

	fmt.Println("Enter the data to put into the file (enter 'exit' on a new line to stop):")
	var text []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)

		if line == "exit" {
			break
		}
	}

	for _, line := range text[:len(text)-1] {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// set the file offset to point at begining
	file.Seek(0, 0)

	fmt.Println("Content read from the file:")
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
