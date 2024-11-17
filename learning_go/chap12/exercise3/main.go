package main

import (
	"fmt"
	"math"
	"sync"
)

func initMap() map[int]float64 {
	results := make(map[int]float64)
	for i := range 100000 {
		results[i] = math.Sqrt(float64(i))
	}
	fmt.Println("Map initialized")
	return results
}

var cachedMap func() map[int]float64 = sync.OnceValue(initMap)

func main() {
	for i := 0; i < 100000; i += 1000 {
		println(cachedMap()[i])
	}

}
