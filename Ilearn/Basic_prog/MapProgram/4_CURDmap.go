// Write a program in to create a stucture of employee with name and sal and perform the CURD on it

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	Emp := make(map[string]int)
	ch := 0
	name, sal := "", 0
	for {
		fmt.Println("---Menu----")
		fmt.Println("1.add employee")
		fmt.Println("2.display employee")
		fmt.Println("3.get Employee")
		fmt.Println("4.delete Employee")
		
		fmt.Println("5.Exit")

		fmt.Println("Enter ur choice :")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter the name :")
			fmt.Scanln(&name)
			fmt.Println("Enter the Salary :")
			fmt.Scanln(&sal)
			Emp[name] = sal
		case 2:
			fmt.Println("--Employee Data---")
			for key, val := range Emp {
				fmt.Println(key, val)
			}
		case 3:

			fmt.Println("Enter the name of employee:")
			fmt.Scanln(&name)
			sal, ishere := Emp[name]
			if ishere {
				fmt.Println(name, sal)
			} else {
				fmt.Println(name, "is not present")
			}
		case 4:

			fmt.Println("Enter the name of employee:")
			fmt.Scanln(&name)
			name = strings.ToLower(name)
			_, ishere := Emp[name]
			if ishere {
				delete(Emp, name)
			}
			fmt.Println("--Employee Data after delete---")
			for key, val := range Emp {
				fmt.Println(key, val)
			}
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")

		}

		fmt.Println()
	}
}
