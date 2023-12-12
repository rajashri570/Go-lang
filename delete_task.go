package main

import "fmt"

func delete_task(id int) {
	tasklist[id].isvalid = false

	fmt.Println("-------after deleting task -----")
	for i := range tasklist {
		if tasklist[i].isvalid == false {
			continue
		} else {
			fmt.Printf("%d \t %s\t%s\t%d %d %s", tasklist[i].Id, tasklist[i].User_name, tasklist[i].Task_name,
				tasklist[i].Status, tasklist[i].Priority, tasklist[i].Deadline)
			fmt.Println()
		}
	}

}
