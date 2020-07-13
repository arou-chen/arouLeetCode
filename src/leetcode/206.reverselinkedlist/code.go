package _06_reverselinkedlist

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for ; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}