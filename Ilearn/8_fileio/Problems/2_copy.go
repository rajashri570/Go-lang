/*
Create a program that copies the content of one file to another.
Allow the user to specify the source and destination file paths.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var file_nm1, file_nm2 string

	//getting path for first file

	fmt.Println("Enter the source file path")
	fmt.Scan(&file_nm1)

	//getting path for second file

	fmt.Println("Enter the destination file path")
	fmt.Scan(&file_nm2)

	//open the source file to get data to copy
	source, err := os.Open(file_nm1)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer source.Close()

	// Get file information
	sourceInfo, err := source.Stat()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Check if the source file is empty
	if sourceInfo.Size() == 0 {
		fmt.Println("Source file is empty. Cannot copy.")
		return
	}

	//creating the destination file
	destination, err := os.Create(file_nm2)

	//coping content of source to destination file
	_, err = io.Copy(destination, source)

	//handling error while coping
	if err != nil {
		fmt.Println("Error in coping file!")
	} else {
		fmt.Println("Coppied successfully")
	}

}
