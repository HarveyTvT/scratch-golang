package main

import (
	"fmt"
	"slices"
)

func test1() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	// fmt.Println(slices.Equal(x, s))

	println(cap(x))
}

func test2() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func test3() {
	x := [4]int{1, 2, 3, 4}
	y := x[:]
	z := [5]int(y) // panic

	fmt.Println(z)
}

func test4() {
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}

func main() {
	test4()
}
