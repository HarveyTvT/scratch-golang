package main

import "fmt"

type LinkedNode[T comparable] struct {
	value T
	next  *LinkedNode[T]
}

type LinkedList[T comparable] struct {
	size int
	head *LinkedNode[T]
	tail *LinkedNode[T]
}

func (l *LinkedList[T]) Add(v T) {
	n := &LinkedNode[T]{value: v}
	l.size++

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
}

func (l *LinkedList[T]) Insert(v T, idx int) {
	l.size++
	n := &LinkedNode[T]{value: v}

	if idx == 0 {
		n.next = l.head
		l.head = n
		return
	}

	cursor := l.head
	for i := 0; i < idx-1; i++ {
		cursor = cursor.next
	}

	n.next = cursor.next
	cursor.next = n
}

func (l *LinkedList[T]) Index(v T) int {
	idx := 0
	cursor := l.head
	for cursor != nil {
		if cursor.value == v {
			return idx
		}
		cursor = cursor.next
		idx++
	}

	return -1
}

func (l *LinkedList[T]) Print() {
	cursor := l.head
	for cursor != nil {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}

func main() {
	n := &LinkedList[int]{}
	n.Add(1)
	n.Add(2)
	n.Add(3)
	n.Insert(0, 2)

	n.Print()
}
