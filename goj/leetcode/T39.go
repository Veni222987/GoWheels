package leetcode

func combinationSum(candidates []int, target int) [][]int {
	// 递归实现
	return nil
}

func cal(candidates []int, target int) (res map[int]int) {
	for _, candidate := range candidates {
		if candidate == target {
			res[candidate]++
		}
	}
	return nil
}
