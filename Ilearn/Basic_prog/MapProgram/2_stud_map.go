//write a program to create a map store the student name and addr and print the student with from pune.

package main

import (
	"fmt"
	"strings"
)

func main() {
	//create map

	StudMap := make(map[int][]string)
	//pune_stud := []string{}

	StudMap[1] = []string{"rajashri", "khed"}
	StudMap[2] = []string{"shubhangi", "pune"}
	StudMap[3] = []string{"kiran", "latur"}
	StudMap[4] = []string{"kishu", "pune"}

	fmt.Println("Students from Pune:")
	for _, val := range StudMap {
		//fmt.Print(key, val[1])
		s := strings.ToLower(val[1])
		if s == "pune" {
			fmt.Println(val[0])
		}

	}

	//fmt.Println(StudMap)

}
