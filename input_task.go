package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func get_user_input(msg string) string {
	var line string

jumpto:
	fmt.Println(msg)

	//used reader to get input using standard input

	reader := bufio.NewReader(os.Stdin)

	//read the text till \n"
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)

		// Return an empty string in case of an error
		return ""
	}

	// Remove the newline character from the end of the line

	line = line[:len(line)-1]
	if reflect.TypeOf(line).Kind() != reflect.String {
		line = ""
		goto jumpto
	}

	return line
}
