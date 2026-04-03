package main

import "fmt"

func main() {
	var input string
	fmt.Printf("Enter a string to reverse : ")
	fmt.Scanf("%v", &input)
	rune_arr := []rune(input)
	fmt.Println(rune_arr)
	for i, j := len(rune_arr)-1, 0; i >= j; i, j = i-1, j+1 {
		rune_arr[i], rune_arr[j] = rune_arr[j], rune_arr[i]
	}
	rev_str := string(rune_arr)
	fmt.Println("Reverse string : ", rev_str)
}
