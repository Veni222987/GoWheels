package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	// 使用O(n)的额外空间解决，n是三角形的行数
	// 实际上只需要使用保存上一层的dp结果就好
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := []int{triangle[0][0]}
	res := math.MaxInt
	for i := 1; i < len(triangle); i++ {
		newLine := make([]int, len(triangle[i]))
		for j := range triangle[i] {
			l, r := math.MaxInt, math.MaxInt
			if j < len(dp) {
				r = dp[j]
			}
			if j-1 >= 0 {
				l = dp[j-1]
			}
			newLine[j] = min(l, r) + triangle[i][j]
			if i == len(triangle)-1 && newLine[j] < res {
				res = newLine[j]
			}
		}
		dp = newLine
	}
	return res
}
