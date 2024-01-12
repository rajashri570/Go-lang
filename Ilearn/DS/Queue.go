// write a program to craate a Queue

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	start *Node
	end   *Node
	size  int
}

func (Q *Queue) Enqueue(data int) {
	node1 := &Node{data: data}

	if Q.size == 0 {
		Q.start = node1
		Q.end = node1
	} else {
		Q.end.next = node1
		Q.end = node1
	}
	Q.size++
	fmt.Println("inserted ele into the queue", data)

}
func (Q *Queue) PrintQueue() {
	if Q.size == 0 {
		fmt.Println("Queue is empty")
		return
	}

	current := Q.start
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}

	fmt.Println()
}

func (Q *Queue) DeQueue() {
	data := 0
	if Q.size == 0 {
		fmt.Println("Queue is empty")
		return
	} else {
		data = Q.start.data
		Q.start = Q.start.next
		Q.size--
	}
	fmt.Print("deleted item :", data)
}

func main() {

	que := Queue{}
	que.PrintQueue()

	que.Enqueue(10)
	que.PrintQueue()
	// fmt.Println(que)

	que.DeQueue()
	que.PrintQueue()
}
