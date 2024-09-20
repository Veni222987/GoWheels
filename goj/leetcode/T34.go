package leetcode

// 在排序数组中查找元素
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r, mid := 0, len(nums)-1, 0
	for l != r && l != r-1 {
		mid = (l + r) >> 1
		if nums[mid] == target {
			break
		} else if nums[(l+r)>>1] < target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[mid] == target {
		l, r = mid, mid
	} else if nums[l] == target {
		r = l
	} else if nums[r] == target {
		l = r
	} else {
		return []int{-1, -1}
	}
	//向两侧扩展
	for ; l-1 >= 0 && nums[l-1] == target; l-- {
	}
	for ; r+1 < len(nums) && nums[r+1] == target; r++ {
	}
	return []int{l, r}
}
