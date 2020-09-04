package _0_pow

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / mul(x, -n)
	}
	return mul(x, n)
}

func mul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := mul(x, n / 2)
	if n % 2 != 0 {
		return result * result * x
	}
	return result * result
}