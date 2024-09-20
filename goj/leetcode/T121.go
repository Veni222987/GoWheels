package leetcode

import "math"

func maxProfit(prices []int) int {
	preMin := make([]int, len(prices))
	preMin[0] = math.MaxInt
	maxPf := math.MinInt
	for i := 1; i < len(preMin); i++ {
		preMin[i] = min(preMin[i-1], prices[i-1])
		maxPf = max(maxPf, prices[i]-preMin[i])
	}
	return max(0, maxPf)
}
