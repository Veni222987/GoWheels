package leetcode

// 搜索插入的位置
// TODO 改造一下，可以变成二分插入的算法
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l != r && l != r-1 {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	if target <= nums[l] {
		return l
	}
	return l + 1
}
