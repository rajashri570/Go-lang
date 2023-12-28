/*
Implement a simple CSV reader and writer program. The reader should take a CSV file as input and output the data in a structured format, while the writer should take structured data and write it to a CSV file.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Emp struct {
	name string
}

func main() {

	file, err := os.Create("emp.csv")
	if err != nil {
		fmt.Println("error in creating file")
	}
	defer file.Close()

	// created 2D slice with interface type so can hold all kind data
	emp := [][]interface{}{
		{1, "rajashri", 25},
		{2, "shiru", 30},
		{3, "sushant", 22},
	}

	for _, data := range emp {

		_, err := fmt.Fprintf(file, "%d , %s , %d\n", data[0], data[1], data[2])
		if err != nil {
			fmt.Println("Error in writing to file")
		}

	}

	file, err = os.Open("emp.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	updated_str := ""
	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, ",")
		if len(slice) == 3 {
			slice[2] = "Pune"
			updated_str += strings.Join(slice, ",") + "\n"
		}
	}

	file, err = os.Create("emp.csv")
	if err != nil {
		fmt.Println("Error in creating file:", err)
		return
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%s", updated_str)
	if err != nil {
		fmt.Println("Error in writing to file:", err)
		return
	}

}
