/*
Build a program that reads a text file, replaces occurrences of a specified word with another word, and then writes the modified content back to the file.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//file, err := os.Open("file2.txt")
	file, err := os.OpenFile("file2.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Errror in opening file")
		return
	}
	defer file.Close()

	var updated_data string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		// fmt.Println(line)
		updated_data = updated_data + strings.Replace(line, "golang", "python", -1) + "\n"
		//updated_data = append(updated_data, strings.Replace(line, "rajashri", "girl", -1))
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(updated_data)

	if err != nil {
		fmt.Println("Error while writing..", err)
		return
	}

	//file.Seek(0, 0)
	err = writer.Flush()
	if err != nil {
		fmt.Println("ERROR OCUURED :", err)
	} else {
		fmt.Println("successfully replaced!")
	}
}
