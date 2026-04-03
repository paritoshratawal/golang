package main

import (
	"fmt"
	"sync"
)

func printCount(count int, wg *sync.WaitGroup, channel <-chan int) {
	defer wg.Done()
	for i := 1; i <= count; i++ {
		fmt.Println(i)
		channel <- i
	}
}

func main() {
	channel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go printCount(10, wg, channel)

	wg.Wait()
	for num := range channel {
		fmt.Println(num)
	}
}
