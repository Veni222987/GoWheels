package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// 初始化行与列
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1]
		if s2[i-1] == s3[dp[0][i-1]] {
			dp[0][i]++
		}
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0]
		if s1[i-1] == s3[dp[i-1][0]] {
			dp[i][0]++
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			num1, num2 := dp[i-1][j], dp[i][j-1]
			if s3[num1] == s1[i-1] {
				num1++
			}
			if s3[num2] == s2[j-1] {
				num2++
			}
			dp[i][j] = max(num1, num2)
		}
	}
	return dp[len(s1)][len(s2)] == len(s3)
}
