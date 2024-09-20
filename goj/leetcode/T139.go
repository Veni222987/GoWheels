package leetcode

// 一维dp即可，但是需要O(n^2)
func wordBreak(s string, wordDict []string) bool {
	checkDict := func(s string) bool {
		for _, v := range wordDict {
			if s == v {
				return true
			}
		}
		return false
	}
	dp := make([]int, len(s))
	for i := range s {
		for j := i; j >= 0; j-- {
			if checkDict(s[dp[j] : i+1]) {
				dp[i] = i + 1
				break
			}
		}
	}
	return dp[len(s)-1] == len(s)
}
