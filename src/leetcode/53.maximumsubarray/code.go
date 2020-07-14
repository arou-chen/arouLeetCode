package _3_maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := math.MinInt64
	step := 0
	for idx, data := range nums {
		if idx == 0 {
			step = data
		} else {
			step = maxInt(step + data, data)
		}
		max = maxInt(max, step)
	}
	return max
}

func maxInt(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
