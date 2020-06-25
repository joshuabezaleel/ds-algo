package stack

import "fmt"

type Item struct {
	val  int
	next *Item
}

type Stack struct {
	top   *Item
	depth int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	item := &Item{val, nil}
	if s.top == nil {
		s.top = item
	} else {
		item.next = s.top
		s.top = item
	}
	s.depth++
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return 0
	}
	top := s.top.val
	s.top = s.top.next
	s.depth--
	return top
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return 0
	}
	return s.top.val
}

func (s *Stack) Print() {
	item := s.top
	for item != nil {
		fmt.Printf("%v ", item.val)
		item = item.next
	}
}
