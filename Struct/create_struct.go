//write a progra to create a  Employee structure and print data of rmployee.
//try auto incremented id

package main

import (
	"bufio"
	"fmt"
	"os"
)

var count = 0

type Salary struct {
	basic int
	HRA   int
}
type Employee struct {
	id   int
	name string
	sal  Salary
}

func getId() int {
	count++
	return count
}
func (emp Employee) getEmp() string {
	return fmt.Sprintf("Id : %d\tName : %s\tSalary: %d\n", emp.id, emp.name, emp.sal.basic+emp.sal.HRA)
}

func main() {
	var n = 0

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the number of employe u want to add : ")
	fmt.Scan(&n)
	var emp = make([]Employee, n)

	for i := 0; i < n; i++ {

		fmt.Println("\nenter the employee name : ")
		emp[i].name, _ = reader.ReadString('\n')
		emp[i].id = getId()
		fmt.Println("enter the salary in terms of basic and hra: ")
		fmt.Scan(&emp[i].sal.basic)
		fmt.Scan(&emp[i].sal.HRA)

	}

	for i := 0; i < n; i++ {
		fmt.Println(emp[i].getEmp())
	}

}
