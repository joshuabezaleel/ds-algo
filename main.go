package main

import (
	"fmt"

	bt "github.com/joshuabezaleel/ds-algo/concepts/ds/binary-tree"
	dll "github.com/joshuabezaleel/ds-algo/concepts/ds/doubly-linked-list"
	queue "github.com/joshuabezaleel/ds-algo/concepts/ds/queue"
	sll "github.com/joshuabezaleel/ds-algo/concepts/ds/singly-linked-list"
	stck "github.com/joshuabezaleel/ds-algo/concepts/ds/stack"
)

func main() {
	singlyLinkedList := sll.NewSinglyLinkedList()
	singlyLinkedList.Append(5)
	singlyLinkedList.Append(3)
	singlyLinkedList.Append(8)
	singlyLinkedList.Append(4)
	singlyLinkedList.Print()
	fmt.Println()

	doublyLinkedList := dll.NewDoublyLinkedList()
	doublyLinkedList.Append(5)
	doublyLinkedList.Append(3)
	doublyLinkedList.Append(8)
	doublyLinkedList.Append(4)
	doublyLinkedList.PrintBackwards()
	fmt.Println()

	singlyLinkedList.Reverse()
	singlyLinkedList.Print()
	fmt.Println()
	singlyLinkedList.Append(6)
	singlyLinkedList.Print()
	fmt.Println()

	fmt.Println("=== STACK ===")
	stack := stck.NewStack()
	stack.Push(5)
	stack.Push(3)
	stack.Push(8)
	stack.Push(4)
	stack.Print()

	fmt.Println("=== QUEUE ===")
	q := queue.NewQueue()
	q.Enqueue(5)
	q.Enqueue(3)
	q.Enqueue(8)
	q.Enqueue(4)
	q.Print()
	fmt.Println()
	fmt.Printf("Dequeue = %v\n", q.Dequeue())
	q.Print()
	fmt.Println()

	fmt.Println("=== BINARY TREE ===")
	bt := bt.NewBinaryTree()
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(8)
	bt.Insert(4)
	bt.PreOrderPrint()
	fmt.Println()
	bt.InOrderPrint()
	fmt.Println()
	bt.PostOrderPrint()
	fmt.Println()
}
