package _47_friendcircles

func findCircleNum(m [][]int) int {
	visitor := make([]bool, len(m))
	time := 0
	for i := 0; i < len(m); i++ {
		if visitor[i] {
			continue
		}
		time++
		dfs(visitor, m, i)
	}
	return time
}

func dfs(visitor []bool, m[][]int, i int) {
	if visitor[i] {
		return
	}
	visitor[i] = true
	for idx, data := range m[i] {
		if data == 1 && idx != i {
			dfs(visitor, m, idx)
		}
	}
}