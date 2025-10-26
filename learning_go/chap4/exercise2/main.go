package main

import "math/rand/v2"

func main() {
	results := make([]int, 0)
	for i := 0; i < 100; i++ {
		results = append(results, rand.IntN(100))
	}

	for i, numb := range results {
		switch {
		case numb%2 == 0 && numb%3 == 0:
			println(i, numb, "Six!")
		case numb%2 == 0:
			println(i, numb, "Two!")
		case numb%3 == 0:
			println(i, numb, "Three!")
		default:
			println(i, numb, "Never mind")
		}
	}
}
