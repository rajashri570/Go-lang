// write a program to sort the map values.

package main

import "fmt"

func main() {

	var MarksMap = make(map[string]int)
	var n int
	fmt.Println("Enter the number of student you want to add :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var name string
		var marks int
		fmt.Println("Enter name of", i+1, "student:")
		fmt.Scanln(&name)
		fmt.Println("Enter Marks of", i+1, "student:")
		fmt.Scanln(&marks)
		MarksMap[name] = marks
	}
	high := 0
	studnm := ""
	for stud, marks := range MarksMap {
		if marks > high {
			high = marks
			studnm = stud
		}
	}

	fmt.Println(studnm, "is having highest marks", high)
}
