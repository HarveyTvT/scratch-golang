package main

import (
	"fmt"
	"time"
)

type Incrementer interface {
	Increment()
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer

	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter    // ok
	myStringer = valueCounter      // ok
	myIncrementer = pointerCounter // ok
	// myIncrementer = valueCounter   // compile-time error!

	fmt.Println(valueCounter.String())

	fmt.Println(myStringer.String())
	myIncrementer.Increment()
}
