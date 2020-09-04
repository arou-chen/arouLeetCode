package __addtwonumber

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:0,
	}
	PResult := result
	var left int
	for {
		var v1, v2 int
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		sum := (v1 + v2 + left) % 10
		left = (v1 + v2 + left) / 10
		PResult.Next = &ListNode{
			Val:sum,
		}
		PResult = PResult.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if left > 0 {
		PResult.Next = &ListNode{
			Val:left,
		}
	}

	return result.Next
}