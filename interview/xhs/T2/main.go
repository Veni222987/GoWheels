package main

import "slices"

func main() {

}

// 排序，使得序列满足中间小两边大，并且大的不能集中在同一边
func twoSideSort(list []int) []int {
	//从大到小排序
	slices.SortFunc(list, func(i, j int) int { return j - i })
	//按照奇偶性分离，并且将其中一个反向拼接到另外一个
	var odd, even []int = make([]int, 0), make([]int, 0)
	for i := range list {
		if i%2 == 0 {
			even = append([]int{list[i]}, even...)
		} else {
			odd = append(odd, list[i])
		}
	}
	odd = append(odd, even...)
	return odd
}
