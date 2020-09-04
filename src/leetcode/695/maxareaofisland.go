package _95_maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	iMax := len(grid)
	if iMax == 0 {
		return 0
	}
	yMax := len(grid[0])
	if yMax == 0 {
		return 0
	}
	max := 0
	for i := 0; i < iMax; i++ {
		for j := 0; j < yMax; j++ {
			if  grid[i][j] == 1 {
				result := dfs(grid, i, j)
				if max < result {
					max = result
				}
			}
		}
	}
	return max
}

func dfs(grid [][]int, i, j int) int {
	iMax := len(grid)
	if i >= iMax || i < 0 {
		return 0
	}
	yMax := len(grid[0])
	if j >= yMax || j < 0 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	result := 1
	result += dfs(grid, i + 1, j)
	result += dfs(grid, i - 1, j)
	result += dfs(grid, i, j + 1)
	result += dfs(grid, i, j - 1)
	return result
}
