package main

import "fmt"

func update_task(id int) {
	var st int
	fmt.Println("enter the status in terms of 0(pending)/1(Completed) : ")
	fmt.Scanln(&st)
	tasklist[id].Status = st

	fmt.Println("---------------Updated Record is -------------")
	fmt.Println()
	fmt.Printf("%d \t %s\t%s\t%d %d %s", tasklist[id].Id, tasklist[id].User_name, tasklist[id].Task_name,
		tasklist[id].Status, tasklist[id].Priority, tasklist[id].Deadline)
	fmt.Println()
}
