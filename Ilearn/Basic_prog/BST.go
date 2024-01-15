package main

import "fmt"

type Node struct {
	data   int
	lchild *Node
	rchild *Node
}

type Btree struct {
	root *Node
	size int
}

func (b *Btree) Insert(data int) {
	newNode := &Node{
		data:   data,
		lchild: nil,
		rchild: nil,
	}

	if b.root == nil {
		b.root = newNode
	} else {
		b.insertNode(b.root, newNode)
	}
	b.size++
}

func (b *Btree) insertNode(root, newNode *Node) {
	if newNode.data < root.data {
		if root.lchild == nil {
			root.lchild = newNode
		} else {
			b.insertNode(root.lchild, newNode)

		}
	} else {
		if root.rchild == nil {
			root.rchild = newNode
		} else {
			b.insertNode(root.rchild, newNode)
			fmt.Println(root.rchild.data)
		}
	}
}

func printPreOrder(node *Node) {
	if node != nil {
		fmt.Print(node.data, " ")
		printPreOrder(node.lchild)

		printPreOrder(node.rchild)
	}
}

func printInOrder(node *Node) {
	if node != nil {
		printInOrder(node.lchild)
		fmt.Print(node.data, " ")
		printInOrder(node.rchild)
	}
}
func printPostOrder(node *Node) {
	if node != nil {
		printPostOrder(node.lchild)
		printPostOrder(node.rchild)
		fmt.Print(node.data, " ")
	}
}

func (B *Btree) SearchKey(key int) {
	if B.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	currentNode := B.root
	lev := 0

	for currentNode != nil {
		{
			if currentNode.data == key {
				fmt.Println("\nKey found at root at level ", lev)
				return
			} else if key < currentNode.data {
				currentNode = currentNode.lchild
			} else {
				currentNode = currentNode.rchild
			}

			//B.size++
			lev++
		}
		//fmt.Println("\nKey not found in tree")
	}
	fmt.Println("\nKey not found in tree")
}

/*
func (B *Btree) deleteKey(key int) {
	if B.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	currentNode := B.root
	lev := 0

	for currentNode != nil {
		{
			if currentNode.data == key {
				fmt.Println(key, "Key deleted at root at level ", lev)
				for currentNode.lchild != nil {
					continue
				}
				currentNode.data = currentNode.lchild.data
				return
			} else if key < currentNode.data {
				currentNode = currentNode.lchild
			} else {
				currentNode = currentNode.rchild
			}

			//B.size++
			lev++
		}
		//fmt.Println("\nKey not found in tree")
	}
	fmt.Println("\nKey not found in tree")

}*/

func main() {
	bt := Btree{}
	bt.Insert(23)
	bt.Insert(15)
	bt.Insert(42)
	bt.Insert(6)
	fmt.Println("InOrder:")
	printInOrder(bt.root)
	fmt.Println("\nPost order:")
	printPostOrder(bt.root)
	fmt.Println("\nPreet order:")
	printPreOrder(bt.root)
	bt.SearchKey(15)
	//t.deleteKey(6)
	//bt.SearchKey(11)

	//fmt.Println(bt.root)
	// You can add more insertions as needed
}
