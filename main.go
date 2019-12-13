package main

import (
	"fmt"
	"projecteuler/problem23"
)

func main() {
	anMap, anSlice := problem23.GetAbundantNumbers(1, 28123)

	//fmt.Println(problem23.HasAbundantAddends(228, &anSlice, &anMap))
	sum := 0
	for i := 1; i < 28123; i++ {
		hasAbundants := problem23.HasAbundantAddends(i, &anSlice, &anMap)
		if !hasAbundants {
			sum += i
		}
	}
	fmt.Println(sum)

}
