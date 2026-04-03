package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is expressive, concise, clean, and efficient. Go is also concurrent and fast"
	new_str := strings.Replace(str, ",", "", -1) //Remove the comma or can remove any character
	fmt.Println("new_str", new_str)
	split_str := strings.Split(new_str, " ")
	occurance := make(map[string]int)
	for _, val := range split_str {
		occurance[val] = occurance[val] + 1
	}
	for key, _ := range occurance {
		if occurance[key] == 1 {
			delete(occurance, key)
		}
	}
	fmt.Println("occurance", occurance)
}
