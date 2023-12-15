//write a program to create a structure of employee and calculate the highest anual income of employe.

package main

import "fmt"

// created struct here

type emp struct {
	id   int
	name string
	addr string
	sal  int
}

func (e emp) getAnualSal() int {
	return e.sal * 12
}

func (e emp) get() string {
	return fmt.Sprintf("Id : %d\tName : %s\tAddres:%s\tAnual sal :%d", e.id, e.name, e.addr, e.sal)
}

func main() {
	e1 := emp{
		id:   1,
		name: "rajashri",
		addr: "khed",
		sal:  20000,
	}
	e2 := emp{
		id:   2,
		name: "suraj",
		addr: "khed",
		sal:  25000,
	}
	fmt.Println(e1.get())
	fmt.Println(e2.get())
	//var high = 0
	fmt.Println()
	if e1.getAnualSal() > e2.getAnualSal() {
		fmt.Println(e1.name, "is having highest anual income of ", e1.getAnualSal())
	} else {
		fmt.Println(e2.name, "is having highest anual income ", e2.getAnualSal())
	}
}
