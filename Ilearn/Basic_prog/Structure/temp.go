//Create a program that uses a struct to represent a student with fields such as name, roll number, and an array to store subject grades.
//Calculate the average grade for each student.

package main

import "fmt"

type Stud struct {
	name  string
	roll  int
	marks [4]int
}

func (s *Stud) PuData() {

	name := ""
	roll := 0
	marks := [4]int{}
	fmt.Println("Enter the name, roll and marks of four sub of Student:")
	fmt.Scan(&name)
	fmt.Scan(&roll)
	for i := 0; i < 4; i++ {
		fmt.Println("Enter ths mark of sub", i+1)
		fmt.Scan(&marks[i])
	}

	s.name = name
	s.roll = roll
	s.marks = marks
}
func (s Stud) GetData() {
	fmt.Println("Name :", s.name, "\nRoll No.:", s.roll, "Marks:", s.marks)
}

func (s Stud) GetAvg() {
	sum := 0.0
	avg := 0.0
	for _, mark := range s.marks {
		sum = sum + float64(mark)
	}
	avg = sum / 4
	fmt.Println("Average marks:", avg)

}
func main() {
	s1 := &Stud{}
	s2 := &Stud{}

	//n := 0

	//marks[0] = 33
	// fmt.Println(marks)
	// fmt.Println("Ente the number os student u want to add :")
	// fmt.Scan(&n)
	// s = make([]Stud, 0, n)

	s1.PuData()

	s1.GetData()
	s1.GetAvg()

	s2.PuData()

	s2.GetData()
	s2.GetAvg()

}
