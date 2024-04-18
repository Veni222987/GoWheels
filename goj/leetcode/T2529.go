package leetcode

func maximumCount(nums []int) int {
	sum, pos := len(nums), 0
	for _, value := range nums {
		if value > 0 {
			pos++
		}
		if value == 0 {
			sum--
		}
	}
	return max(pos, sum-pos)
}
