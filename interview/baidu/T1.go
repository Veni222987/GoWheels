package main

func allSubset(input []string) [][]string {
	// 回溯
	var backtrack func(rest []string) [][]string
	backtrack = func(rest []string) [][]string {
		if len(rest) == 1 {
			return [][]string{{rest[0]}, {""}}
		}
		res := make([][]string, 0)
		for i := 0; i < len(rest)-1; i++ {
			nextRes := backtrack(rest[i+1:])
			// 子集+当前元素
			for j := range nextRes {
				var tmp []string
				for k := range nextRes[j] {
					if nextRes[j][k] != "" {
						tmp = append([]string{rest[i]}, nextRes[j][k])
					}
				}
				res = append(res, tmp)
			}
			res = append(res, nextRes...)
			// 当前元素
			res = append(res, []string{rest[i]})
		}
		return res
	}
	// 统一处理空子集
	return backtrack(input)
}
