package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	str := "ababbaabb"
	splt_str := strings.Split(str, "")
	// fmt.Println(splt_str)
	fmt.Println(slices.Compact(splt_str))
}
