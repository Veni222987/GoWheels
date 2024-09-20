package leetcode

func subsets79(nums []int) [][]int {
	// 回溯
	var backtrack func(rest []int) [][]int
	backtrack = func(rest []int) [][]int {
		if len(rest) == 1 {
			return [][]int{{rest[0]}, make([]int, 0)}
		}
		res := make([][]int, 0)
		nextRes := backtrack(rest[1:])
		// 子集+当前元素
		for j := range nextRes {
			tmp := append([]int{rest[0]}, nextRes[j]...)
			res = append(res, tmp)
		}
		res = append(res, nextRes...)
		return res
	}
	return backtrack(nums)
}
