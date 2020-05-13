package _02_binarytreelevelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		temp := make([]int, 0)
		nodeQ := make([]*TreeNode, 0)
		for _, data := range q {
			temp = append(temp, data.Val)
			if data.Left != nil {
				nodeQ = append(nodeQ, data.Left)
			}
			if data.Right != nil {
				nodeQ = append(nodeQ, data.Right)
			}
		}
		q = nodeQ
		ret = append(ret, temp)
	}
	return ret
}
