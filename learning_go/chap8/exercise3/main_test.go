package main

import "testing"

func TestMain(m *testing.T) {
	n := &LinkedList[int]{}
	n.Add(1)
	n.Add(2)
	n.Add(3)
	n.Insert(0, 0)

	n.Print()
}
