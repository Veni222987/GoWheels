package leetcode

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	vHead := &ListNode{}
	tail := vHead
	for {
		end := true
		minNode := &ListNode{Val: math.MaxInt}
		mIndex := 0
		for i := range lists {
			if lists[i] != nil {
				end = false
				if lists[i].Val <= minNode.Val {
					minNode = lists[i]
					mIndex = i
				}
			}
		}
		if end {
			break
		}
		tail.Next = minNode
		tail = tail.Next
		lists[mIndex] = lists[mIndex].Next
	}
	return vHead.Next
}
