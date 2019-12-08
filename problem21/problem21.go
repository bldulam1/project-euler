package problem21

import "math"

// Divisors returns all the divisors of a number
func Divisors(n int) []int {
	max := int(math.Sqrt(float64(n)))
	divisors := make([]int, 0)

	for i := 1; i <= max; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if n/i != i && i > 1 {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

// Sum returns the sum of an array
func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// AmicablePair checks if the number is amicable
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair
func AmicablePair(num int) int {
	b := Sum(Divisors(num))

	if num == Sum(Divisors(b)) {
		return b
	}

	return -1

}
