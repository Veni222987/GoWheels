package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	disL, disR := 0, 0
	for i := len(distance) - 1; i >= destination; i-- {
		disL += distance[i%len(distance)]
	}
	for i := 0; i <= destination; i++ {
		disR += distance[i]
	}
	return min(disL, disR)
}
