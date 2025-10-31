package main

func main() {
	var i float64

	println(i / 0) // NaN

	i = 1
	println(i / 0) // +Inf

	i = -1
	println(i / 0) // -Inf
}
