package leetcode

import "slices"

// 组合总和
func combinationSum(candidates []int, target int) [][]int {
	//遍历candidates
	//为了确保有序性，这里只取比当前数小的
	// 为了减少遍历次数，这里先对candidate排序
	slices.Sort(candidates)

	var backtrack func(sortedCandidates []int, target, upperBound int) [][]int
	backtrack = func(sortedCandidates []int, target, upperBound int) [][]int {
		res := make([][]int, 0)
		for _, v := range sortedCandidates {
			if v > target || v > upperBound {
				break
			}
			if v < target {
				nextRes := backtrack(sortedCandidates, target-v, v)
				for i := range nextRes {
					nextRes[i] = append(nextRes[i], v)
				}
				res = append(res, nextRes...)
			}
			if v == target {
				res = append(res, []int{v})
			}
		}
		return res
	}
	return backtrack(candidates, target, target)
}
