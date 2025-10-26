package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make([]int, 0)
	for i := 0; i < 100; i++ {
		results = append(results, rand.IntN(100))
	}

	fmt.Println(results)
	fmt.Println(len(results))
}
