package utils

// Min returns the minimum
func Min(nums []int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// Max returns the maximum from the array
func Max(nums []int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
