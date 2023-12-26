//write a program to print the last 100 bytes from file

package main

import (
	"fmt"
	"os"
)

func main() {
	var filenm string
	//open the file
	fmt.Println("Enter the filename :")
	fmt.Scan(&filenm)

	file, err := os.Open(filenm)

	if err != nil {
		fmt.Println("ERROR OCCURED : ")
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat() //it creates offset to hold the info of file like nm,size..
	if err != nil {
		fmt.Println("Error for getting file info", err)
		return
	}
	start_offset := fileInfo.Size() - 100
	_, err = file.Seek(start_offset, 0) //it make file offset set to 100's byte from last bydefault it 0,0

	if err != nil {
		fmt.Println("Error in seek file")
	}

	buffer := make([]byte, 100) // here we have just created cutome buffersize
	_, err = file.Read(buffer)  // reading bytes set for buffer
	if err != nil {
		fmt.Println("Error in reding buffer data")
	}
	fmt.Println("bytes from string :\n", string(buffer)) // conversion required to get the strings but
	//if not convetrd then its gives byte values for character.
}
