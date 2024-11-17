package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// change GOGC and GODEBUG to see the difference
func main() {
	startedAt := time.Now()
	results := make([]Person, 0, 1000000)
	// results := make([]Person, 0)

	results = append(results, Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	})

	fmt.Println(len(results))
	fmt.Println("Time taken: ", time.Since(startedAt), " seconds")
}
