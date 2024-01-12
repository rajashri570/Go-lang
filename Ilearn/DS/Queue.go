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

	if Q.size == 0 {
		fmt.Println("Queue is empty")
		return
	} else {
		data := Q.end.data
		q.end.next = q.end - 1
	}
}

func main() {

	que := Queue{}
	que.PrintQueue()

	que.Enqueue(10)
	que.PrintQueue()
	// fmt.Println(que)
}
