package main

import "fmt"

func display() {
	for i := range tasklist {
		fmt.Println() //added bcos ater getting choice it put extra line
		fmt.Printf("%d \t %s\t%s\t%d %d %s", tasklist[i].Id, tasklist[i].User_name, tasklist[i].Task_name,
			tasklist[i].Status, tasklist[i].Priority, tasklist[i].Deadline)
		fmt.Println()
	}
}
