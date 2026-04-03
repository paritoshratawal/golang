package main

import "fmt"

func main(){
	arr := []int{1, 2, 3, 4}
	contantTimeComplexity(arr)
	linearTimeComplexity(arr)
	unsortedArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", unsortedArr)
	sorted := bubbleSort(append([]int{}, unsortedArr...))
	fmt.Println("Bubble sorted array:", sorted)
	mergeSorted := mergeSort(unsortedArr)
	fmt.Println("Merge sorted array:", mergeSorted)
}

//BIG O NOTATION: O(1) Constant Time
func contantTimeComplexity(arr []int){
	fmt.Println("constatnt time complexity")
	fmt.Println(arr[0])
}

//BIG O NOTATION: O(n) Linear Time
func linearTimeComplexity(arr []int){
	fmt.Println("linear time complexity")
	for _, v := range arr{
		fmt.Println(v)
	}

}

//BIG O NOTATION: O(n^2) Quadratic Time
func bubbleSort(arr []int) []int{
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++{
		for j := 0; j < arrLen-i-1; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}			
		}
	}
	return arr
}

//BIG O NOTATION: O(n log n) Linearithmic Time
func mergeSort(arr []int) []int {
	i := 0
	i++
	fmt.Println("merge sort call:", i)
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	fmt.Println("merging")
	fmt.Println("left:", left, "right:", right)
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// //BIG O NOTATION: O(2^n) Exponential Time
// func fibonacci(n int) int{

// }

// //BIG O NOTATION: O(n!) Factorial Time
// func travelingSalesmanProblem(graph [][]int) int{

// }	

// //BIG O NOTATION: O(n^k) Polynomial Time
// func polynomialTimeComplexity(arr []int) []int{

// }

// //BIG O NOTATION: O(k^n) Exponential Time with base k
// func exponentialTimeComplexity(n int) int{

// }