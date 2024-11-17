package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range 10 {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := range 10 {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		if ch1 == nil && ch2 == nil {
			break
		}
		select {
		case i, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(i)
			}
		case i, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(i)
			}
		}

	}
}
