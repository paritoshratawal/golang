package main

import "fmt"

type myArray struct {
	total_size int
	used_size  int
	array_ptr  []int
}

func createArray(array *myArray, total_students int, enrolled_students int) {
	array.total_size = total_students
	array.used_size = enrolled_students
	array.array_ptr = make([]int, total_students)
}

func insertValues(array *myArray) {
	for i := 0; i < array.used_size; i++ {
		fmt.Printf("Enter marks for student %d:", i+1)
		fmt.Scanln(&array.array_ptr[i])
	}
}

func show(array *myArray) {
	for i := 0; i < array.used_size; i++ {
		fmt.Printf("Marks of student %d: %d\n", i+1, array.array_ptr[i])
	}
}

func main() {
	marks := myArray{}
	total_students := 100
	enrolled_students := 5
	createArray(&marks, total_students, enrolled_students)
	insertValues(&marks)
	show(&marks)
	// fmt.Println(marks)
}
