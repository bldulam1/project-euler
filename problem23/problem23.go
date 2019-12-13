package problem23

import (
	"projecteuler/problem21"
)

// Abundance checks if the number is abundant, 1 if abundant, 0 if perfect, and -1 if deficient
func Abundance(number int) int8 {
	divisors := problem21.Divisors(number)
	if problem21.Sum(divisors) < number {
		return -1
	} else if problem21.Sum(divisors) > number {
		return 1
	}
	return 0
}

func GetAbundantNumbers(min int, max int) (map[int]bool, []int) {
	anMap := make(map[int]bool)
	anSlice := make([]int, 0)

	for i := min; i < max; i++ {
		if Abundance(i) > 0 {
			anMap[i] = true
			anSlice = append(anSlice, i)
		}
	}

	return anMap, anSlice
}

func HasAbundantAddends(num int, addends *[]int, addMap *map[int]bool) bool {
	var add1, add2 int

	addsLen := len(*addends)

	for i1 := 0; i1 < addsLen && (*addends)[i1] < num; i1++ {
		add1 = (*addends)[i1]
		add2 = num - add1
		if (*addMap)[add2] {
			return true
		}
	}
	return false
}
