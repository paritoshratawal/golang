package main

import "fmt"

func generate_odd(count int, odd_ch chan<- int) {
	for i := 1; i < count; i++ {
		if i%2 != 0 {
			odd_ch <- i
		}
	}
	close(odd_ch)
}

func generate_even(count int, even_ch chan<- int) {
	for i := 1; i < count; i++ {
		if i%2 == 0 {
			even_ch <- i
		}
	}
	close(even_ch)
}
func main() {
	even_ch := make(chan int)
	odd_ch := make(chan int)

	go generate_even(10, even_ch)
	go generate_odd(10, odd_ch)

	for odd := range odd_ch {
		fmt.Println("Odd", odd)
	}

	for even := range even_ch {
		fmt.Println("Even", even)
	}

}
