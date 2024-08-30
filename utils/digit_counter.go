package utils

import "math"

// ! 目标:统计某一个数字的位数

// ! 1、当数据不会出现0的时候，可以使用以下方法统计
func DigitCounter(num int) int {
	return int(math.Log10(float64(num))) + 1
}
