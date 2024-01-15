package main

type Node1 struct {
	data   int
	lchild *Node
	rchild *Node
}

type Btree1 struct {
	root *Node
	size int
}

func (b *Btree1) Insert(data int) {
	newNode1 := &Node1{
		data:   data,
		lchild: nil,
		rchild: nil,
	}

	if b.root == nil {
		b.root = newNode1
	} else {
		b.insertNode(b.root, newNode)
	}
	b.size++
}

func (b *Btree1) insertNode(root, newNode *Node1) {
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
		}
	}
}

func main() {
	bt := Btree1{}
	bt.Insert(23)
	bt.Insert(15)
	bt.Insert(42)

	// You can add more insertions as needed
}
