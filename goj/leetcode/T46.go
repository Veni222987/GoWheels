package leetcode

// 数字全排列
// 主要思想：回溯
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	res := make([][]int, 0)
	// 选择一个数字，其余的回溯
	for i, num := range nums {
		restArr := make([]int, len(nums)-1)
		copy(restArr, nums[:i])
		copy(restArr[i:], nums[i+1:])
		for _, arr := range permute(restArr) {
			res = append(res, append(arr, num))
		}
	}
	return res
}

// 题解
func permute2(nums []int) [][]int {
	l := len(nums)
	res := [][]int{}
	// ! 想要回溯不能快速赋值
	var backtrack func(int)
	backtrack = func(first int) {
		if first == l {
			res = append(res, append([]int(nil), nums...))
			return
		}
		for i := first; i < l; i++ {
			// 通过换位实现了遍历的功能
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
}
