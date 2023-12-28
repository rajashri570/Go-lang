/*
Write a program that generates a tree-like structure representing the hierarchy of files and directories starting from a specified root directory. Use os and os.FileInfo to gather information about files and directories.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//open file
	file, err := os.Open("file1.txt")

	//error handling
	if err != nil {
		fmt.Println("ERROR OCCURED :", err)
	}
	defer file.Close()
	fmt.Println()

	//geting the current directory
	cd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error in printing curent directory..")
		return
	}

	//preparing the file path
	filePath := cd + "/" + file.Name()
	fmt.Println("File path is :", filePath)

	// seperating the names from full path
	pathslice := strings.Split(filePath, "/")
	//fmt.Print(pathslice)

	// printng path with indentaion
	for i, name := range pathslice {
		indentation := strings.Repeat(" ", i)
		fmt.Printf("%s|__%s\n", indentation, name)
	}

}
