//write a program to delete if the file is empty otherwise display file text

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var file_nm string
	fmt.Println("Enther the filename with extension :")
	fmt.Scan(&file_nm)

	fileInfo, err := os.Stat(file_nm)

	if err != nil {
		if os.IsNotExist(err) {

			log.Fatal("File does not exist")

		} else {

			fmt.Print("file exist..")
		}

	} else {

		if fileInfo.Size() == 0 {
			fmt.Println("file is empty")
			os.Remove(file_nm)
		} else {
			data, err := ioutil.ReadFile(file_nm)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("--------------------------")
				fmt.Println(string(data))
			}

		}

	}
}
