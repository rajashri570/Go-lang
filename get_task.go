package main

import "fmt"

func get_task(id int) {
	// fmt.Println("taskid :", tasklist[0].Id)
	fmt.Println("Name :", tasklist[id].User_name)
	fmt.Println("Description :", tasklist[id].Task_name)
	fmt.Println("Status :", tasklist[id].Status)
	fmt.Println("Priority :", tasklist[id].Priority)
	fmt.Println("Deadline :", tasklist[id].Deadline)
}
