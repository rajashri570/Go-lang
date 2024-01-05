package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type List struct {
	head *Node
}

// AddNode adds a new node with the given data to the end of the list.
func (l *List) AddNode(d string) {
	newNode := &Node{data: d}

	if l.head == nil {
		// If the list is empty, set the new node as the head
		l.head = newNode
	} else {
		// Traverse the list to find the last node and add the new node
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// PrintList prints the data of all nodes in the list.
func (l *List) PrintList() {
	count := 0
	//nodeMap := make(map[int]*Node)
	//len = 0
	//isVisited = false
	current := l.head
	for current != nil {
		fmt.Println("Node Data:", current.data)
		current = current.next
		count++
		//len = count
	}
	if current == nil {
		fmt.Print("No cycle")
	}

}

func (l *List) IsCycle() {

}

func main() {
	// Create a linked list
	myList := &List{}

	// Add nodes to the list
	myList.AddNode("rajashri")
	myList.AddNode("Suraj")
	myList.AddNode("shera")
	myList.AddNode("Shrikant")

	// Print the list
	fmt.Println("Linked List Contents:")
	myList.PrintList()
}
