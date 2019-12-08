package main

import (
	"fmt"
	"projecteuler/problem20"
)

func main() {
	f100 := problem20.BigFactorial(100)
	sum := 0
	for _, c := range f100 {

		if c > 48 {
			sum += int(c) - 48
		}
	}
	fmt.Println(sum)

}
