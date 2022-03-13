package week3
type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    if n == 1 {
        return lists[0]
    }
    if n == 2 {
        return mergeTwoListNode(lists[0],lists[1])
    }
    l1 := mergeKLists(lists[:n/2])
    l2 := mergeKLists(lists[n/2:])
    return mergeTwoListNode(l1,l2)
}

func mergeTwoListNode(l1,l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    prev := dummy
    for l1 != nil || l2 != nil {
        if l2 == nil ||  (l1 != nil && l1.Val < l2.Val) {
            prev.Next = l1
            l1 = l1.Next
        }else {
            prev.Next = l2
            l2 = l2.Next
        }
        prev = prev.Next
    }
    return dummy.Next
}

