package dynamicprogramming

func maximalSquare(matrix [][]byte) int {
	var max int
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			dp[i] = append(dp[i], int(matrix[i][j] - '0'))
			if dp[i][j] == 1 {
				max = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if dp[i][j] == 0 {
				continue
			}
			val := min(min(dp[i][j-1], dp[i-1][j]), dp[i - 1][j - 1]) + 1
			dp[i][j] = val
			if val > max {
				max = val
			}
		}
	}

	return max * max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}