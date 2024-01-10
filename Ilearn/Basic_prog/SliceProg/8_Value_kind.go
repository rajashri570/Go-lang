package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Example with an integer
	// var num interface{}
	// fmt.Println("Enter the number :")
	// fmt.Scan(&num)
	// value := reflect.ValueOf(num)

	// // Print the type and value
	// fmt.Printf("Type: %s\n", value.Type())
	// fmt.Printf("Value: %v\n", value.Interface())

	str := "Hello, Go!"
	value := reflect.ValueOf(str)

	// Print the type and value
	fmt.Printf("Type: %s\n", value.Type())
	fmt.Printf("Value: %v\n", value.Interface())

}
