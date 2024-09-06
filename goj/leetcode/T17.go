package leetcode

// 电话号码的字母组合(这跟电话号码有半毛钱关系)
func letterCombinations(digits string) []string {
	numMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "yuv",
		'9': "wxyz",
	}
	strArr := []string{}
	for _, v := range digits {
		strArr = append(strArr, numMap[v])
	}
	res := make([]string, 0)
	for _, str := range strArr {
		if len(res) == 0 {
			for i := range str {
				res = append(res, str[i:i+1])
			}
		} else {
			l := len(res)
			for i := 0; i < l; i++ {
				for j := range str {
					res = append(res, res[i]+str[j:j+1])
				}
			}
			res = res[l:]
		}
	}
	return res
}
