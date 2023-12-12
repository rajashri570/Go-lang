package main

import (
	"fmt"
	"reflect"
	"time"
)

func add_task() {
	unm := ""
	tnm := ""
	st := 0
	pr := 0
	dl := 0
	var prompt string

	prompt = "Enter the User name : "
	unm = get_user_input(prompt)

	prompt = "Enter the task name :"
	tnm = get_user_input(prompt)

	for {
		fmt.Println("Enter the Priority Task (1-3):")
		_, err := fmt.Scanln(&pr)
		if err == nil && reflect.TypeOf(pr).Kind() == reflect.Int {
			break
		} else {
			fmt.Println("Invalid input for Priority. Please enter an integer.")
		}
	}

	for {
		fmt.Println("Enter the deadline count in days :")
		_, err := fmt.Scanln(&dl)
		if err == nil && reflect.TypeOf(dl).Kind() == reflect.Int {
			break
		} else {
			fmt.Println("Invalid input for Deadline. Please enter an integer.")
		}
	}

	date := time.Now().AddDate(0, 0, dl)

	tsk := Task{
		Id:        len(tasklist),
		User_name: unm,
		Task_name: tnm,
		Status:    st,
		Priority:  pr,
		Deadline:  date,
		isvalid:   true,
	}

	tasklist = append(tasklist, tsk)

}
