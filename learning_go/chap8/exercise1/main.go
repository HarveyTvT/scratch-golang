package main

import "fmt"

func Double[T interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}](v T) T {
	return 2 * v
}

func main() {
	fmt.Println(Double(8))
	fmt.Println(Double(int8(8)))
	fmt.Println(Double(8.8))
}
