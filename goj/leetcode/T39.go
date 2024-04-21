package leetcode

func combinationSum(candidates []int, target int) [][]int {
	// 递归实现
	m := make(map[int]int)
	res := make([][]int, 0)

	return res
}

func cal(candidates []int, target int) (res map[int]int) {
	for _, candidate := range candidates {
		if candidate == target {
			res[candidate]++
		}
	}
}
