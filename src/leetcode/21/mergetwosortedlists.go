package _1_mergetwosortedlists

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode
	result := l
	for  {
		if l1 == nil {
			l = l2
			break
		}
		if l2 == nil {
			l = l1
			break
		}
		if l1.Val <= l2.Val {
			l = l1
			l1 = l1.Next
		}
	}

	return result
}
