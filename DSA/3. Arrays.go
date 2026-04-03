package main

import (
	"errors"
	"fmt"
)

func indexInsertValues(arr [100]int, size int, element int, index int) {
	fmt.Println("Size of array before insertion", size)
	capacity := len(arr)
	if size >= capacity {
		fmt.Errorf("Cannot insert element : ", errors.New("Array capacity exceeded"))
	}
	for i := size - 1; i >= index; i-- {
		arr[i+1] = arr[i]
	}
	arr[index] = element
	size++
	display(arr, size)
}

func display(arr [100]int, size int) {
	fmt.Println("Size of array after insertion", size)
	for i := 0; i < size; i++ {
		fmt.Println("Element", arr[i])
	}
}

func main() {
	arr := [100]int{1, 3, 5, 9, 22}
	size := 5
	element_to_be_insert := 2
	index_to_be_insert := 1
	indexInsertValues(arr, size, element_to_be_insert, index_to_be_insert)
}
