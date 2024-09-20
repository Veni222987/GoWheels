package leetcode

// ! 无法解决中位数的两个数位于同一个数组的情况
// TODO 感觉还可以优化，有空可以来挑战一下
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return (float64((nums2[len(nums2)/2])) + float64(nums2[(len(nums2)-1)/2])) / 2
	}
	if len(nums2) == 0 {
		return (float64((nums1[len(nums1)/2])) + float64(nums1[(len(nums1)-1)/2])) / 2
	}

	totalLen := len(nums1) + len(nums2)
	l1, r1, l2, r2 := 0, len(nums1)-1, 0, len(nums2)-1
	m1, m2 := (l1+r1)/2, (l2+r2)/2
	// 比较两个m求和
	for m1+m2 != (totalLen-1)/2 {
		m1, m2 = (l1+r1)/2, (l2+r2)/2
		if m1+m2 < (totalLen-1)/2 {
			// 将小的改大
			if nums1[m1] < nums2[m2] && l1 < len(nums1)-1 {
				l1 = m1 + 1
			} else {
				l2 = m2 + 1
			}
		} else {
			// 将大的改小
			if nums1[m1] > nums2[m2] && r1 > 0 {
				r1 = m1 - 1
			} else {
				r2 = m2 - 1
			}
		}
	}
	if totalLen%2 == 0 {
		return (float64(nums1[m1]) + float64(nums2[m2])) / 2
	}
	return float64(min(nums1[m1], nums2[m2]))
}
