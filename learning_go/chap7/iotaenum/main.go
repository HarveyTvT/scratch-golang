package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = 1 << iota
	Personal
	Spam
	Social
	Advertisements
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

func main() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
}
