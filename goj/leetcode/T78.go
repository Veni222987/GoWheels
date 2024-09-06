package leetcode

// T78 子集
// ! 这里一开始踩坑了，必须用两层深拷贝再append，不然扩容节点之后会出问题
func subsets(nums []int) [][]int {
	var backtrack func(first int) [][]int
	backtrack = func(first int) [][]int {
		res := [][]int{{}}
		if first == len(nums) {
			return res
		}
		tmp := backtrack(first + 1)
		res = make([][]int, len(tmp))
		for i, v := range tmp {
			res[i] = make([]int, len(v))
			copy(res[i], v)
		}
		for _, v := range tmp {
			res = append(res, append(v, nums[first]))
		}
		return res
	}
	return backtrack(0)
}
