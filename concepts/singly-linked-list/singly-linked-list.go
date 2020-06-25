package singlylinkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	length int
	head   *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sll *SinglyLinkedList) Append(value int) {
	node := &Node{value, nil}
	if sll.head == nil {
		sll.head = node
	} else {
		last := sll.head
		for last.next != nil {
			last = last.next
		}
		last.next = node
	}
	sll.length++
}

func (sll *SinglyLinkedList) Print() {
	node := sll.head
	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
}

func (sll *SinglyLinkedList) Reverse() {
	var prev *Node = nil
	var curr = sll.head
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	sll.head = prev
}
