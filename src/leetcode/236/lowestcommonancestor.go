package _36_lowestcommonancestor

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if right != nil {
		return right
	}
	return left
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := map[int]*TreeNode{}
	m := map[int]bool{}
	var save func (node *TreeNode)
	save = func (node *TreeNode) {
		if node.Left != nil {
			parent[node.Left.Val] = node
			save(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			save(node.Right)
		}
	}
	save(root)

	for p != nil {
		m[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if m[q.Val] {
			return q
		}
		m[q.Val] = true
		q = parent[q.Val]
	}

	return nil
}