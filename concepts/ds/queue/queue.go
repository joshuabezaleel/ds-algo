package queue

import "fmt"

type Item struct {
	val  int
	next *Item
}

type Queue struct {
	front  *Item
	length int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	item := &Item{val, nil}
	if q.front == nil {
		q.front = item
	} else {
		last := q.front
		for last.next != nil {
			last = last.next
		}
		last.next = item
	}
	q.length++
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		return 0
	}
	frontVal := q.front.val
	q.front = q.front.next
	q.length--
	return frontVal
}

func (q *Queue) Peek() int {
	if q.front == nil {
		return 0
	}
	return q.front.val
}

func (q *Queue) Print() {
	item := q.front
	for item != nil {
		fmt.Printf("%v ", item.val)
		item = item.next
	}
}
