package main

/*
   #cgo LDFLAGS: -lm
   #include "mini_calc.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	x := '+'
	f := (*C.char)(unsafe.Pointer(&x))
	fmt.Println(C.mini_calc(f, 1, 2))
}
