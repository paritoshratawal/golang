package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	//Here we need to create a array of string
	split_str := strings.Split(str, "")
	for i, j := len(split_str)-1, 0; i >= j; i, j = i-1, j+1 {
		split_str[i], split_str[j] = split_str[j], split_str[i]
	}
	rev_str := strings.Join(split_str, "")

	fmt.Println(split_str)
	fmt.Println(rev_str)
}
