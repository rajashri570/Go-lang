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

	file, err := os.Open("file2.txt")
	if err != nil {
		fmt.Println("Errror in opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	var updated_data string

	for scanner.Scan() {
		
		line := scanner.Text()

		updated_data = updated_data + strings.Replace(line, "golang", "python",-1)+"\n"
	}
	file.Seek(0,0)
	writer := bufio.NewWriter(file)
	_,err = writer.WriteString(updated_data)

	if err != nil{
		fmt.Println("Error while writing..",err)
		return
	} else{
		fmt.Println("data replaced successfully")
	}
	err = writer.Flush()
	if err != nil{
		fmt.Println("error in flush")
		return
	}

}


		
		/*line = strings.TrimRight(line, "\n")
		wordList := strings.Fields(line)

		for i, word := range wordList {

			if word == "rajashri" {
				fmt.Println("index:", i)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("ERROR READING words:", err)
			return
		}*/



