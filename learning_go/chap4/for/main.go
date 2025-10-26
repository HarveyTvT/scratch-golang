package main

import "fmt"

func test1() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 3; i++ {
		for k, v := range m {
			println(i, k, v)
		}
	}
}

func test2() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func test3() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func main() {
	test3()
}
