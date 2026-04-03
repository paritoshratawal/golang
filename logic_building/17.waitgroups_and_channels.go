package main

import (
	"fmt"
	"sync"
)

func print_counting(wg *sync.WaitGroup, ch chan<- int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}

	wg.Done()
	// close(ch)
}
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go print_counting(wg, ch)
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(ch) // Close the channel
	}()
	for num := range ch {
		fmt.Println("num", num)
	}
}
