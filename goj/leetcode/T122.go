package leetcode

import "math"

func maxProfit2(prices []int) int {
	preMin := make([]int, len(prices))
	preMin[0] = math.MaxInt
	pf := 0
	for i := 1; i < len(preMin); i++ {
		preMin[i] = min(preMin[i-1], prices[i-1])
		if prices[i] > preMin[i] {
			pf += prices[i] - preMin[i]
			preMin[i] = math.MaxInt
		}
	}
	return pf
}
