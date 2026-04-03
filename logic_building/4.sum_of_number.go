package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number:")
	fmt.Scanf("%v", &a)
	sum := 0
	for a > 0 {
		sum = sum + a%10
		a = a / 10
		fmt.Println("a", a)
	}
	fmt.Println("sum of a number:", sum)
}
