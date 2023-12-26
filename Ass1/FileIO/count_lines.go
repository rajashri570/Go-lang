// write a program to open the file and count the no of line in file.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var file *os.File
	var count = 0
	file, err := os.Open("file_create.txt")
	if err != nil {
		fmt.Print("Error: ", err)
	} else {

		reader := bufio.NewScanner(file)
		for reader.Scan() {
			count += 1
			//fmt.Println(reader.Text())
		}

	}
	defer file.Close()
	fmt.Println("number of lines in file is :", count)
}
