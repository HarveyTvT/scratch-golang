package main

import "fmt"

func failUpdate(g *int) {
	f := 10
	g = &f
}

func main() {
	var f *int
	failUpdate(f)
	fmt.Println(f)
}
