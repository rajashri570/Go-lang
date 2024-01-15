// package main

// type Node struct {
// 	data   int
// 	lchild *Node
// 	rchild *Node
// }

// type Btree struct {
// 	root *Node
// 	size int
// }

// func (b *Btree) Insert(data int) {
// 	newNode := &Node{
// 		data:   data,
// 		lchild: nil,
// 		rchild: nil,
// 	}

// 	if b.root == nil {
// 		b.root = newNode
// 	} else {
// 		b.insertNode(b.root, newNode)
// 	}
// 	b.size++
// }

// func (b *Btree) insertNode(root, newNode *Node) {
// 	if newNode.data < root.data {
// 		if root.lchild == nil {
// 			root.lchild = newNode
// 		} else {
// 			b.insertNode(root.lchild, newNode)
// 		}
// 	} else {
// 		if root.rchild == nil {
// 			root.rchild = newNode
// 		} else {
// 			b.insertNode(root.rchild, newNode)
// 		}
// 	}
// }

// func main() {
// 	bt := Btree{}
// 	bt.Insert(23)
// 	bt.Insert(15)
// 	bt.Insert(42)

// 	// You can add more insertions as needed
// }

//write w progream to create a binary search tree.

package main

import "fmt"

type NodeStruct struct {
	data   int
	lchild *NodeStruct
	rchild *NodeStruct
}

type Btree struct {
	root *NodeStruct
	size int
}

func (b *Btree) Insert(data int) {

	node1 := &NodeStruct{
		data:   data,
		lchild: nil,
		rchild: nil,
	}
	if b.size == 0 {
		b.root = node1
	} else if b.size == 1 {
		if node1.data < b.root.data {
			b.root.lchild = node1
			//b.root = b.root.lchild
			fmt.Println(node1.data)
		} else {
			b.root.rchild = node1
			//b.root = b.root.rchild
			fmt.Println(node1.data)
		}
	} else {
		fmt.Println("not done")
	}
	if b.root.lchild == nil && b.root.rchild == nil {
		b.size++
	}

}

func main() {
	bt := &Btree{}
	bt.Insert(23)
	bt.Insert(11)
	bt.Insert(91)
	fmt.Println(bt.root)
}
