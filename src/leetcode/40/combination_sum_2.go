package _0

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	comb := make([]int, 0)
	dfs := func(idx, target int) {}
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
		isEnd := true
		nextId := idx
		for nextId++; nextId < len(candidates); nextId++ {
			nextNum := candidates[nextId]
			if nextNum != curNum {
				isEnd = false
				break
			}
		}
		if !isEnd {
			dfs(nextId, target)
		}
		if delta > 0 {
			comb = append(comb, curNum)
			dfs(idx + 1, delta)
			comb = comb[:len(comb) - 1]
		}
	}
	dfs(0, target)
	return result
}