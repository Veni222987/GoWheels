package leetcode

import "slices"

func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	m := n / 2
	res := 0
	for i, j := 0, m; i < m && j < n; i++ {
		for j < n && 2*nums[i] > nums[j] {
			j++
		}
		if j < n {
			res += 2
			j++
		}
	}
	return res
}
