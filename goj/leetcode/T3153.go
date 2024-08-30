package leetcode

func sumDigitDifferences(nums []int) int64 {
	type ds struct {
		counter int
		m       map[int]int
	}
	len := 0
	for temp := nums[0]; temp > 0; temp /= 10 {
		len++
	}
	mapArr := make([]ds, len)
	for i := 0; i < int(len); i++ {
		mapArr[i] = ds{
			counter: 0,
			m:       make(map[int]int),
		}
	}
	// 统计数位map
	sum := 0
	for _, num := range nums {
		tmpNum := num
		for i := 0; i < len; i++ {
			// 统计数位
			sum += (mapArr[i].counter - mapArr[i].m[tmpNum%10])
			mapArr[i].counter++
			mapArr[i].m[tmpNum%10]++
			tmpNum /= 10
		}
	}
	return int64(sum)
}
