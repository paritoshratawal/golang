package main

import "fmt"

func reverse(str_runes []rune) {
	for i := len(str_runes) - 1; i > 0; i-- {
		fmt.Println(str_runes[i])
		// rune_chan <- str_runes[i]
	}
	// close(rune_chan)
}

func main() {
	str := "aaDDDwwRRRRwwDDDaa"
	rune_str := []rune(str)
	str_chan := make(chan string)
	reverse(rune_str, str_chan)
	// for rune_val := range rune_chan {
	// fmt.Println(string(rune_val))
	// }
}
