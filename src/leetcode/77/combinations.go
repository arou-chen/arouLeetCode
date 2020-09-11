package _7

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	comb := make([]int, 0)
	dfs := func(idx int) {}
	dfs = func(idx int) {
		if idx > n || len(comb) >= k {
			return
		}
		comb = append(comb, idx)
		if len(comb) == k {
			result = append(result, append([]int{}, comb...))
		}
		dfs(idx+1)
		comb = comb[:len(comb)-1]
		dfs(idx+1)
	}
	dfs(1)
	return result
}