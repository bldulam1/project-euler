package main

import (
	"fmt"
	"projecteuler/problem17"
	"strings"
)

func main() {
	sum := 0
	for i := uint16(1); i <= 1000; i++ {
		word := problem17.NumberToWord(i)
		word = strings.Replace(word, " ", "", -1)
		word = strings.Replace(word, "-", "", -1)
		// fmt.Println(word)
		sum += len(word)
	}
	fmt.Println(sum)

}
