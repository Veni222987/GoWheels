package leetcode

import (
	"fmt"
	"testing"
)

func Test222(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{4, nil, nil},
			Right: &TreeNode{5, nil, nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{6, nil, nil},
			Right: nil,
		},
	}
	fmt.Println(countNodes(&root))
}
