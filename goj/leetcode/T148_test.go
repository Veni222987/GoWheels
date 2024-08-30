package leetcode

import (
	"fmt"
	"testing"
)

func TestT148(t *testing.T) {
	input := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}

	for head := sortList(input); head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
