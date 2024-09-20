package week415

import "math"

func maxScore(a []int, b []int) int64 {
	dp := make([][]int64, len(a))
	for i := range dp {
		dp[i] = make([]int64, len(b))
	}
	// 初始化第一行
	for i := range len(b) {
		if i >= 1 {
			dp[0][i] = max(dp[0][i-1], int64(a[0])*int64(b[i]))
		} else {
			dp[0][i] = int64(a[0]) * int64(b[i])
		}
	}
	// 开始DP
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if j < i {
				dp[i][j] = math.MinInt
				continue
			}
			// 当前值匹配
			chooseValue := int64(a[i])*int64(b[j]) + dp[i-1][j-1]
			dp[i][j] = max(chooseValue, dp[i][j-1])
		}
	}
	return dp[len(a)-1][len(b)-1]
}
