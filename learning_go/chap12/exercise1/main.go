package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		done <- struct{}{}
	}()

	wg.Wait()
	close(ch)

	<-done
}
