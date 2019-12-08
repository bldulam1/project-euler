package main

import (
	"fmt"
	"projecteuler/problem16"
)

func main() {

	// fmt.Println(problem16.BigProduct("1099511627776", "1099511627776", 5))
	pow := problem16.BigPower("2", 1000)

	sumDigits := uint64(0)

	for i := range pow {
		sumDigits += uint64(pow[i] - 48)
	}

	fmt.Println(sumDigits)
}
