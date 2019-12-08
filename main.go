package main

import (
	"fmt"
	"projecteuler/problem21"
)

func main() {

	pairs := make(map[int]int, 0)
	sum := 0
	for i := 1; i < 10000; i++ {
		if pairs[i] > 0 {
			continue
		}

		pair := problem21.AmicablePair(i)

		if pair >= 0 && i != pair {
			sum += pair + i
			pairs[pair] = i
		}
	}
	fmt.Println(sum)

}
