package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go long_process(done)
	time.Sleep(3 * time.Second)
	close(done)
}

func long_process(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Do Work!!")
		}
	}
}
