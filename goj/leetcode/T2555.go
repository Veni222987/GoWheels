package leetcode

import "sort"

// 两个线段获得的最多奖品
func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	dp := make([]int, n+1)
	ans := 0
	for i := 0; i < n; i++ {
		x := sort.SearchInts(prizePositions, prizePositions[i]-k)
		ans = max(ans, i-x+1+dp[x])
		dp[i+1] = max(dp[i], i-x+1)
	}
	return ans
}
