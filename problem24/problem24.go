package problem24

import (
	"fmt"
	"sort"
	"strings"
)

// Sort the string in lexicographically
// ascending order
func NPermute(str string, n uint64) {
	chars := GetChars(str)
	fmt.Println(strings.Join(chars, ""))

	i := uint64(0)
	for i < n-1 {
		NextPermutation(chars)
		i++
	}

	//2783915460
	fmt.Println(i, strings.Join(chars, ""))
}

// GetMap returns the map character: count
// Example "AABC" => {A:2, B:1, C:1}
func GetChars(str string) []string {
	ms := make([]string, 0)
	for _, i2 := range str {
		ms = append(ms, string(i2))
	}
	sort.Strings(ms)
	return ms
}

func NextPermutation(p []string) bool {
	for a := len(p) - 2; a >= 0; a-- {
		if p[a] < p[a+1] {
			for b := len(p) - 1; ; b-- {
				if p[b] > p[a] {
					//swap
					Swap(p, a, b)

					a++

					for b = len(p) - 1; a < b; b-- {
						Swap(p, a, b)
						a++
					}
					return true
				}
			}
		}
	}

	return false
}

func Swap(p []string, a int, b int) {
	t := p[a]
	p[a] = p[b]
	p[b] = t
}
