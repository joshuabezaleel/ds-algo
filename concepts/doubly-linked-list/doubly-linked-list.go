package doublylinkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Append(value int) {
	node := &Node{value, nil, nil}
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		last := dll.head
		for last.next != nil {
			last = last.next
		}
		last.next = node
		node.prev = last
		dll.tail = node
	}
	dll.length++
}

func (dll *DoublyLinkedList) PrintBackwards() {
	node := dll.tail
	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.prev
	}
}
