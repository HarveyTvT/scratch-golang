package main

import "fmt"

func test1() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}

func main() {
	test1()
}
