package leetcode

func numberOfPoints(nums [][]int) int {
	points := make(map[int]bool)
	for _, car := range nums {
		start, end := car[0], car[1]
		for i := start; i <= end; i++ {
			points[i] = true
		}
	}
	count := len(points)
	return count
}
