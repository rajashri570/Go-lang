//Create a program that uses a struct to represent a student with fields such as name, roll number, and an array to store subject grades.
//Calculate the average grade for each student.

package main

import "fmt"

type Stud struct {
	name  string
	roll  int
	marks [4]int
}

func (s *Stud) PuData(nm string, roll int, marks [4]int) {

	s.name = nm
	s.roll = roll
	s.marks = marks
}
func (s Stud) GetData() {
	fmt.Println("Name :", s.name, "\nRoll No.:", s.roll, "Marks:", s.marks)
}
func main() {
	//s := []Stud{}
	var s []Stud
	// s[0].name = "raj"
	// s[0].roll = 1
	// s[0].marks = [4]int{23, 33, 44, 55}
	// fmt.Println(s)

	n := 0
	name := ""
	roll := 0
	marks := [4]int{}
	//marks[0] = 33
	fmt.Println(marks)
	fmt.Println("Ente the number os student u want to add :")
	fmt.Scan(&n)
	s = make([]Stud, 0, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter the name, roll and marks of four sub of Student:")
		fmt.Scan(&name)
		fmt.Scan(&roll)
		for i := 0; i < 4; i++ {
			fmt.Println("Enter ths mark of sub", i+1)
			fmt.Scan(&marks[i])
		}

		s = []Stud{}
		s[i].PuData(name, roll, marks)

	}

	for _, data := range s {
		data.GetData()

	}
}
