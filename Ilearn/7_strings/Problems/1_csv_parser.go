/*
Create a CSV parser that takes a CSV string as input and converts it into a two-dimensional slice or map. Use the strings package to split the input into rows and columns.
*/
package main

//importing required packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//parsing comma seperatd values to  map
func getCSV(str string) (map[int]string, error) {

	//dividing the string into slice of words
	word_slice := strings.Split(str, ",")

	// fmt.Print((word_slice))

	//created map with name csv_map
	var csv_map = make(map[int]string)

	//putting value from slice to map
	for i, val := range word_slice {
		csv_map[i] = val
	}

	return csv_map, nil
}
func main() {

	var input_str string
	reader := bufio.NewReader(os.Stdin)

	//get string string with comma seperated values form user
	fmt.Println("Enther the string as comma seperated values :")

	input_str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occured!")
		return
	}
	//parse csv
	CSV_Parser, err := getCSV(input_str)
	if err != nil {
		fmt.Println("ERROR in parsing valuels:", err)
		return
	}
	for key, val := range CSV_Parser {
		fmt.Printf("CSV_Parser[%d]: %s\n", key, val)
	}

}
