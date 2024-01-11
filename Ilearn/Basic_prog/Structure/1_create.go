//write a program to create a struct and print the struct values

package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	name   string
	salary int
	addr   string
}

func (s *Employee) Put_data(name string, sal int, add string) {
	s.name = name
	s.salary = sal
	s.addr = add
}
func (s Employee) Disp_data() {
	fmt.Println(s.name, "lives in ", s.addr, "and is having salary", s.salary)
}
func main() {
	s1 := new(Employee)
	fmt.Println(reflect.ValueOf(s1).Type())
	s1.Put_data("rajashri", 30000, "shiroli")
	s1.addr = "Latur"
	s1.Disp_data()

	s2 := Employee{
		name:   "suraj",
		salary: 32000,
		addr:   "Pune",
	}
	s2.Disp_data()
}
