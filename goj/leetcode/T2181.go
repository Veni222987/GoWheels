package leetcode

// 合并零之间的节点
func mergeNodes(head *ListNode) *ListNode {
	virtualHead := &ListNode{}
	tail := virtualHead
	var sum int
	for ; head != nil; head = head.Next {
		if head.Val == 0 {
			if sum != 0 {
				tail.Next = &ListNode{sum, nil}
				tail = tail.Next
				sum = 0
			}
		} else {
			sum += head.Val
		}
	}
	return virtualHead.Next
}
