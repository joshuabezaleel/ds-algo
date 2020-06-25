package binarytree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(val int) {
	if bt.root == nil {
		bt.root = &Node{val, nil, nil}
	} else {
		bt.root.Insert(val)
	}
}

func (n *Node) Insert(val int) {
	if val <= n.val {
		if n.left == nil {
			n.left = &Node{val, nil, nil}
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{val, nil, nil}
		} else {
			n.right.Insert(val)
		}
	}
}

func (bt *BinaryTree) PreOrderPrint() {
	preOrderPrint(bt.root)
}

func preOrderPrint(n *Node) {
	if n != nil {
		fmt.Printf("%v ", n.val)
		preOrderPrint(n.left)
		preOrderPrint(n.right)
	}
}

func (bt *BinaryTree) InOrderPrint() {
	inOrderPrint(bt.root)
}

func inOrderPrint(n *Node) {
	if n != nil {
		inOrderPrint(n.left)
		fmt.Printf("%v ", n.val)
		inOrderPrint(n.right)
	}
}

func (bt *BinaryTree) PostOrderPrint() {
	postOrderPrint(bt.root)
}

func postOrderPrint(n *Node) {
	if n != nil {
		postOrderPrint(n.left)
		postOrderPrint(n.right)
		fmt.Printf("%v ", n.val)
	}
}
