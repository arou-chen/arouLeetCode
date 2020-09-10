package _9

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	comb := make([]int, 0)
	var dfs func(idx, target int)
	dfs = func(idx, target int) {
		if idx >= len(candidates) {
			return
		}
		curNum := candidates[idx]
		delta := target - curNum
		if delta == 0 {
			comb = append(comb, curNum)
			result = append(result, append([]int{}, comb...))
			comb = comb[:len(comb) - 1]
		}
		dfs(idx+1, target)
		if delta > 0 {
			comb = append(comb, curNum)
			dfs(idx, delta)
			comb = comb[:len(comb) - 1]
		}
	}
	dfs(0, target)
	return result
}
