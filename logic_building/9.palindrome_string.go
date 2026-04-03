package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Please enter a string to check palindrome:")
	fmt.Scanf("%v", &input)
	str_array := strings.Split(input, "")
	for i, j := len(str_array)-1, 0; i > j; i, j = i-1, j+1 {
		str_array[i], str_array[j] = str_array[j], str_array[i]
	}
	rev_str := strings.Join(str_array, "")
	if input == rev_str {
		fmt.Println(input, "is palindrome")
	} else {
		fmt.Println(input, "is not palindrome")
	}
}
