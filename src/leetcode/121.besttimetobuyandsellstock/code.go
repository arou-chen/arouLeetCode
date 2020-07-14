package _21_besttimetobuyandsellstock

import "math"

func maxProfit(prices []int) int {
	maxDelta := 0
	minVal := math.MaxInt64
	if len(prices) == 0 {
		return 0
	}
	for _, data := range prices {
		if data < minVal {
			minVal = data
		}
		delta := data - minVal
		if maxDelta < delta {
			maxDelta = delta
		}
	}
	return maxDelta
}
