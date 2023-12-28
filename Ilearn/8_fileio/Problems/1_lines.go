/*
Write a program that reads a text file and counts the number of lines in it. Use bufio.Reader for efficient line-by-line reading.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(writer *bufio.Writer, data string) error {

	_, err := writer.WriteString(data)

	if err != nil {
		fmt.Println("Error in writing!")

	}
	return writer.Flush()
}

func count_lines(reader *bufio.Reader) int {

	count := 0

	for {
		_, err := reader.ReadString('\n')

		count++

		if err != nil {
			break // End of input or error
		}

	}
	//fmt.Println(count)
	return count
}

func main() {

	//creating file
	file, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println("Error in file creation!")
		return
	}
	defer file.Close()

	//creating buffer for writing

	writer := bufio.NewWriter(file)

	data := "I am rajashri\nI am learning Go. Its great programming language."

	//writing to file

	err = write(writer, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	//positioning the poiter at begining of file
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	line_count := count_lines(reader)

	if line_count == 0 {
		fmt.Print("file is Empty!")
	}

	fmt.Println("No of lines in", file.Name(), line_count)

}
