package leetcode

import "math"

func findMin(nums []int) int {
	minNum := math.MaxInt
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[l] < nums[m] {
			// mid的左侧是有序的
			minNum = min(minNum, nums[l], nums[r])
			l = m + 1
		} else {
			minNum = min(minNum, nums[m], nums[r])
			r = m - 1
		}
	}
	return minNum
}
