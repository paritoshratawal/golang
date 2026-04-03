package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	score := []int{0}

	wg.Add(3) //3 goroutine to be added
	go func(wg *sync.WaitGroup) {
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg)

	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
