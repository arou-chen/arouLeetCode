package __longestsubstring

func LengthOfLongestSubstring(s string) int {
	length := len(s)
	idx, max := 0, 0
	for i := 0; i < length; i++ {
		step := 0
		idx = i
		m := map[byte]int{}
		for ; idx < length && m[s[idx]] == 0; idx++ {
			step++
			m[s[idx]] = 1
		}
		if max < step {
			max = step
		}
	}
	return max
}