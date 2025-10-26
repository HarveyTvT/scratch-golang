package main

import "math"

func main() {
	var (
		b      byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	println(b, b+1)
	println(smallI, smallI+1)
	println(bigI, bigI+1)
}
