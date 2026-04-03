package main

import "fmt"

func generate_counting(count int, ch chan<- int) {
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	count := 10
	ch := make(chan int)

	go generate_counting(count, ch)
	for num := range ch { // Loop until the channel is closed
		fmt.Println(num)
	}

}
