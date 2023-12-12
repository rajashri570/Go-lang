package main

import (
	//"bufio"
	"fmt"
	"os"
	"time"
)

var count = 0

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
}

//
func add_task() {
	/*reader := bufio.NewReader(os.Stdin)
	for {
		// ReadString('\n') reads until the next newline character
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading standard input:", err)
			break
		}

		// TrimSpace removes leading and trailing whitespaces, including the newline character
		line = bufio.TrimSpace(line)

		if line == "exit" {
			break
		}
	}*/
	unm := ""
	tnm := ""
	st := 0
	pr := 0
	dl := 0
	fmt.Println("Enter the username : ")
	fmt.Scanln(&unm)
	fmt.Println("Enter the task name :")
	fmt.Scanln(&tnm)
	fmt.Println("Enter the Priority Task :")
	fmt.Scanln(&pr)
	fmt.Println("Enter the deadline count in days :")
	fmt.Scanln(&dl)
	date := time.Now().AddDate(0, 0, dl)

	tsk := Task{
		Id:        len(tasklist),
		User_name: unm,
		Task_name: tnm,
		Status:    st,
		Priority:  pr,
		Deadline:  date,
	}

	tasklist = append(tasklist, tsk)

}

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
func get_task(id int) {
	// fmt.Println("taskid :", tasklist[0].Id)
	fmt.Println("Name :", tasklist[id].User_name)
	fmt.Println("Description :", tasklist[id].Task_name)
	fmt.Println("Status :", tasklist[id].Status)
	fmt.Println("Priority :", tasklist[id].Priority)
	fmt.Println("Deadline :", tasklist[id].Deadline)
}
func display() {
	for i := range tasklist {
		fmt.Println() //added bcos ater getting choice it put extra line
		fmt.Printf("%d \t %s\t%s\t%d %d %s", tasklist[i].Id, tasklist[i].User_name, tasklist[i].Task_name,
			tasklist[i].Status, tasklist[i].Priority, tasklist[i].Deadline)
		fmt.Println()
	}
}

func main() {

	var ch int
	for {
		fmt.Println("------Menu---------")
		fmt.Println("1. Add Task")
		fmt.Println("2.Get Task")
		fmt.Println("3.Update Task")
		fmt.Println("4.Delete Task")
		fmt.Println("5.List Tasks")
		fmt.Println("6.Quit")
		fmt.Println("------------------")
		fmt.Print("Enter Your choice : - ")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			add_task()
		case 2:
			var id int
			fmt.Println("enter the task id to display :")
			fmt.Scan(&id)
			get_task(id)
		case 3:
			var id int
			fmt.Println("enter the task id to Update :")
			fmt.Scan(&id)
			update_task(id)
		case 4:
			fmt.Print("empty")
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
