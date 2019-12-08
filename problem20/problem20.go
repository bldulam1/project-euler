package problem20

import (
	"projecteuler/problem16"
)

// BigFactorial performs factorial in strings
func BigFactorial(num int) string {
	acc := "1"
	n := "1"
	for i := 2; i <= num; i++ {
		n = problem16.BigSum([]string{n, "1"})
		acc = problem16.BigProduct(acc, n)
		// println(i+1, acc)
	}
	return acc
}

func factorial(num int) int {
	if num < 2 {
		return 1
	}

	return num * factorial(num-1)
}
