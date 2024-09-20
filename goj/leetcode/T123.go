package leetcode

import "math"

func maxProfit3(prices []int) int {
	preMin := make([]int, len(prices))
	preMin[0] = math.MaxInt
	pf := []int{}
	for i := 1; i < len(preMin); i++ {
		preMin[i] = min(preMin[i-1], prices[i-1])
		if prices[i] > preMin[i] {
			if len(pf) < 2 {
				pf = append(pf, prices[i]-preMin[i])
			} else {
				if prices[i]-preMin[i] > min(pf[0], pf[1]) {
					pf = []int{max(pf[0], pf[1]), prices[i] - preMin[i]}
				}
			}
			preMin[i] = math.MaxInt
		}
	}
	res := 0
	for i := 0; i < len(pf); i++ {
		res += max(0, pf[i])
	}
	return res
}
