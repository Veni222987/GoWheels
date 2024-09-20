package week415

func getSneakyNumbers(nums []int) []int {
	set := make(map[int]byte)
	res := make([]int, 0, 2)
	for _, v := range nums {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
		set[v] = '1'
	}
	return res
}
