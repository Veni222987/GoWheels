package leetcode

// dp求解方法
func getRow(rowIndex int) []int {
	dp := []int{1}
	for i := 1; i <= rowIndex; i++ {
		newLine := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				newLine[j] = 1
				continue
			}
			newLine[j] = dp[j-1] + dp[j]
		}
		dp = newLine
	}
	return dp
}
