package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter the three numbers to find the largest number:")
	fmt.Scanf("%v%v%v", &a, &b, &c)
	lrgst := a
	if lrgst < b {
		lrgst = b
	}
	if lrgst < c {
		lrgst = c
	}

	fmt.Println(lrgst, "is the largest")
}
