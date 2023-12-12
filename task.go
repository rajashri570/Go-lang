package main

import (
	//"bufio"

	"fmt"
	"os"
	"time"
)

//var count int (? - did not get the increatmented value in Task struct to add in slice.

const (
	//STARTED   = "TODO"
	Pending   = 0
	Completed = 1
	High      = 1
	Medium    = 2
	Low       = 3
)

var tasklist []Task

type Task struct {
	Id        int
	User_name string
	Task_name string
	Status    int
	Priority  int
	Deadline  time.Time
	isvalid   bool
}

// func isValidName(name string) bool {
// 	for _, char := range name {
// 		if unicode.IsDigit(char) {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {

	var ch int
	var id int
	for {
		fmt.Println("------Menu---------")
		fmt.Println("1. Add Task")
		fmt.Println("2. Get Task")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. List Tasks")
		fmt.Println("6. Quit")
		fmt.Println("---------------------")
		fmt.Print("Enter Your choice : - ")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			add_task()
		case 2:
			//var id int
			fmt.Println("enter the task id to display :")
			fmt.Scan(&id)
			get_task(id)
		case 3:

			fmt.Println("enter the task id to Update :")
			fmt.Scan(&id)
			update_task(id)
		case 4:
			fmt.Println("enter the task id to delete: ")
			fmt.Scan(&id)
			delete_task(id)
		case 5:
			display()
		case 6:
			os.Exit(0)
		}
		fmt.Println("Press Enter to continue : ")
		fmt.Scanln()
	}
}

//Query  : time.now().Format("2006-01-02")
/*  Task to complete today
1) have to use bufio package
2) delete the slice elem




*/
