package sqrt

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	l, r := 0, x
	result := 0

	for l <= r {
		mid := l + (r - l) / 2
		sqrtVal := mid * mid
		if sqrtVal <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}