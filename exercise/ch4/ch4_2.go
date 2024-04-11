package ch4

/*
针对Unicode的回文串判断

此处考察string类型的编码方式，注意len(str)是字节数不是字符数
*/
func isPalindrome(str string) bool {
	runeSlice := make([]rune, 0)
	for _, u := range str {
		runeSlice = append(runeSlice, u)
	}
	for i, runeLen := 0, len(runeSlice); i < (runeLen+1)/2; i++ {
		if runeSlice[i] != runeSlice[runeLen-1-i] {
			return false
		}
	}
	return true
}
