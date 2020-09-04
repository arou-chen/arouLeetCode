package _12_pathsum

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root.Val <= sum {
		delta := sum - root.Val
		var lRes bool
		var rRes bool
		if root.Left != nil {
			lRes = hasPathSum(root.Left, delta)
		}
		if root.Right != nil {
			rRes = hasPathSum(root.Right, delta)
		}
		return lRes || rRes
	} else {
		return false
	}
}
