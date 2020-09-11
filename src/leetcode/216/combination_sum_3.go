package _16

func combinationSum3(k int, n int) [][]int {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := make([][]int, 0)
	comb := make([]int, 0)
	dfs := func(idx, target int){}
	dfs = func(idx, target int) {
		if idx >= 9 || len(comb) >= k {
			return
		}
		curNum := list[idx]
		delta := target - curNum
		if delta == 0 {
			comb = append(comb, curNum)
			if len(comb) == k {
				result = append(result, append([]int{}, comb...))
			}
			comb = comb[:len(comb)-1]
		}
		dfs(idx+1, target)
		if delta > 0 {
			comb = append(comb, curNum)
			dfs(idx + 1, delta)
			comb = comb[:len(comb) - 1]
		}
	}
	dfs(0, n)
	return result
}