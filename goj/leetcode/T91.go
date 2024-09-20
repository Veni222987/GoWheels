package leetcode

import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '0' {
		return 0
	} else {
		dp[1] = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '0' {
				return 0
			} else {
				dp[i+1] = dp[i-1]
			}
		} else {
			// 和前面的组合与26比较
			num, _ := strconv.Atoi(s[i-1 : i+1])
			if num < 10 {
				dp[i+1] = dp[i]
			} else if num <= 26 && num > 10 {
				dp[i+1] = dp[i]
				if i >= 1 {
					dp[i+1] += dp[i-1]
				}
			} else {
				dp[i+1] = dp[i]
			}
		}
	}
	return dp[len(s)]
}

func numDecodingsOfficial(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}
