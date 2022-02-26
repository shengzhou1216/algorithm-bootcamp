package week1

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	for p1, p2, prev := list1, list2, dummy; p1 != nil || p2 != nil; {
		// 何时添加p1？ p2 为空 或 p1的值小于p2时
		if p2 == nil || (p1 != nil && p1.Val < p2.Val) {
			prev.Next = p1
			p1 = p1.Next
		} else {
			prev.Next = p2
			p2 = p2.Next
		}
		prev = prev.Next
	}
	return dummy.Next
}
