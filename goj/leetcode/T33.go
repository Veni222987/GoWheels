package leetcode

func search(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	//判断位于哪一段
	for l != r && l != r-1 {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}
		mid = (l + r) >> 1
		if nums[l] < nums[mid] {
			//l到mid是有序的
			if target > nums[l] && target < nums[mid] {
				r = mid
			} else {
				l = mid
			}
		} else {
			if target > nums[mid] && target < nums[r] {
				l = mid
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
