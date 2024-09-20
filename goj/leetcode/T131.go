package leetcode

func partition(s string) [][]string {
	// 一维dp就可以了，保存从开始到当前值的子串有多少个
	//题目保证len(s)>=1
	dp := make([][][]string, len(s))
	dp[0] = [][]string{{s[0:1]}}
	for i := 1; i < len(s); i++ {
		for j := 0; j < i+1; j++ {
			if isReverse(s[j : i+1]) {
				//深拷贝原来的串
				if j == 0 {
					dp[i] = [][]string{{s[0 : i+1]}}
				} else {
					for k := range dp[j-1] {
						tmp := append([]string{}, dp[j-1][k]...)
						dp[i] = append(dp[i], append(tmp, s[j:i+1]))
					}
				}
			}
		}
	}
	return dp[len(s)-1]
}

// isReverse
func isReverse(s string) bool {
	l, r := 0, len(s)-1
	for ; l != r && l != r-1; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	if l == r-1 && s[l] != s[r] {
		return false
	}
	return true
}
