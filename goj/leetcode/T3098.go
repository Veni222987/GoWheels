package leetcode

import "math"

/*
1. 找出所有长度为k的子序列, O(n!)
2. 子序列长度最小值，先排序，再遍历，O(nlogn+n)
*/

func sumOfPowers(nums []int, k int) int {
	min := math.MaxInt
	sortedSubArrs := findSubSeq(nums, k, 0)
	for _, arr := range sortedSubArrs {
		// k>=2
		for i := 1; i < len(arr); i++ {
			if absInt(arr[i]-arr[i-1]) < min {
				min = absInt(arr[i] - arr[i-1])
			}
		}
	}
	return min
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 从index开始（包含index）找k个长度的子序列
func findSubSeq(nums []int, k, index int) [][]int {
	if k == 1 {
		rsp := make([][]int, len(nums)-index)
		for i := index; i < len(nums); i++ {
			rsp = append(rsp, []int{nums[i]})
		}
		return rsp
	}
	rsp := findSubSeq(nums, k-1, index+1)
	for _, v := range rsp {
		// 原来的子序列都是有序的，直接用二分插入
		binInsert(v, nums[index])
	}
	return rsp
}

// 二分插入
func binInsert(arr []int, num int) []int {
	left, right, mid := 0, len(arr), 0
	for left != right && left != right-1 {
		mid = (left + right) / 2
		if num == arr[mid] {
			break
		} else if num > arr[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	tmp := append([]int{num}, arr[mid+1:]...)
	return append(arr[0:mid+1], tmp...)
}
