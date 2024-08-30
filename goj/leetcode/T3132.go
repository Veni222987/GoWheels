package leetcode

import (
	"slices"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	//差值栈
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
