package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter a number:")
	fmt.Scanf("%v", &num)
	rev := 0
	copy := num
	for num > 0 {
		rev = rev*10 + num%10
		num = num / 10
	}

	if copy == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
