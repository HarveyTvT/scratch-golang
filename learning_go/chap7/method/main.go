package main

import (
	"fmt"
)

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func test1() {
	myAdder := Adder{start: 10}
	f1 := myAdder.AddTo
	fmt.Println(f1(10))
}

func test2() {
	f := Adder.AddTo
	fmt.Println(f(Adder{10}, 10))
}

func main() {
	test1()

	test2()
}
