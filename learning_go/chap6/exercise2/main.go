package main

import "fmt"

func UpdateSlice(slice []string, value string) {
	if len(slice) > 0 {
		slice[len(slice)-1] = value
		fmt.Println(slice)
	}
}

func GrowSlice(slice []string, value string) {
	slice = append(slice, value)
	fmt.Println(slice)
}

func main() {
	slice := make([]string, 3, 5)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println(cap(slice))
	UpdateSlice(slice, "d")
	fmt.Println(slice)
	GrowSlice(slice, "e")
	fmt.Println(slice)
}
