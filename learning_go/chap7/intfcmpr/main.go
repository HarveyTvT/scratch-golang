package main

import (
	"fmt"
	"reflect"
)

type Doubler interface {
	Double()
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func DoubleCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func test1() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	dis := DoubleIntSlice{1, 2, 3}
	dis2 := DoubleIntSlice{1, 2, 3}

	DoubleCompare(&di, &di2)
	DoubleCompare(&di, dis)

	// DoubleCompare(dis, dis2) // panic
	if v := reflect.ValueOf(dis2); v.Comparable() {
		DoubleCompare(dis, dis2)
	}
}

func main() {
	test1()
}
