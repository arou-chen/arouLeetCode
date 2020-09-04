package _1

func solveNQueens(n int) [][]string {
	verticalList := make([]bool, n)
	front := make(map[int]bool)
	backList := make(map[int]bool)
	for i := n; i < n; i++ {
		verticalList[i] = false
	}

}