package _57

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pathList []string

func binaryTreePaths(root *TreeNode) []string {
	pathList = make([]string, 0)
	if root == nil {
		return pathList
	}
	dfs(root, "")
	return pathList
}

func dfs(node *TreeNode, path string) {
	if node != nil {
		path += strconv.Itoa(node.Val)
		if node.Right == nil && node.Left == nil {
			pathList = append(pathList, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
}