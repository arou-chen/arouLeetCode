package _20_triangle

import "math"

func minimumTotal(triangle [][]int) int {
	iLen := len(triangle)
	if iLen == 0 || len(triangle[0]) == 0 {
		return 0
	}
	temp := make([]int, iLen)
	temp[0] = triangle[0][0]
	for i := 1; i < iLen; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				temp[j] = temp[j] + triangle[i][j]
			} else if i == j {
				temp[j] = temp[j - 1] + triangle[i][j]
			} else {
				temp[j] = min(temp[j-1], temp[j]) + triangle[i][j]
			}
		}
	}
	minVal := math.MaxInt64
	for _, data := range temp {
		minVal = min(minVal, data)
	}
	return minVal
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}