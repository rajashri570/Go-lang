//write a program to create a stack and perform the basic operations.

package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (st *Stack) Push(data int) {
	node1 := &Node{
		data: data,
		next: st.top,
	}
	st.top = node1
	fmt.Println("data pushed successfully")
	//fmt.Println(st.top)

}
func (st Stack) Pop() int {
	if st.IsEmpty() {
		return -1
	} else {
		data := st.top.data
		st.top = st.top.next
		return data

	}
}
func (st Stack) IsEmpty() bool {
	if st.top == nil {
		return true
	} else {
		return false
	}

}
func main() {

	var st = Stack{}
	var ch int
	var data int
	for {
		fmt.Println("----Menu-------")
		fmt.Println("1.push")
		fmt.Println("2.Pop")
		fmt.Println("3.Check For Empty")
		fmt.Println("4.Exit")

		fmt.Println("\nEnter your choice :")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter the data to push:")
			fmt.Scan(&data)
			st.Push(data)
		case 2:
			ele := st.Pop()
			fmt.Println("Poped element :", ele)
		case 3:
			if st.IsEmpty() {
				fmt.Println("Stack is Empty")
			} else {
				fmt.Println("stack is not empty")
			}

		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid Choice!")

		}
	}
}
