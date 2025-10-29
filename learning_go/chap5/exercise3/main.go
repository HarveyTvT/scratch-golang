package main

import "fmt"

func prefixer(in string) func(string) string {
	return func(s string) string {
		return in + s
	}
}

func main() {
	helloPrefix := prefixer("Hello ")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
