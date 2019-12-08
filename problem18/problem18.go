package problem18

// MaxPathSum returns the minimum cost path of a matrix
func MaxPathSum(tri [][]int, m int, n int) int {
	// loop for bottom-up calculation
	for i := m - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			// for each element, check both
			// elements just below the number
			// and below right to the number
			// add the maximum of them to it
			if tri[i+1][j] > tri[i+1][j+1] {
				tri[i][j] += tri[i+1][j]
			} else {
				tri[i][j] += tri[i+1][j+1]
			}
		}
	}

	return tri[0][0]
}
