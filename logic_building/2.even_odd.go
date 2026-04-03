package main

import "fmt"

func even_odd_check(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	var num int
	fmt.Println("Enter the number to check Even or Odd:")
	fmt.Scanf("%v", &num)
	if even_odd_check(num) {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd number")
	}
}
