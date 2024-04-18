package leetcode

import "strings"

func isValidSerialization(preorder string) bool {
	strArr := strings.Split(preorder, ",")
	// 因为有#，可以直接当成完全二叉树，使用数组遍历
	numCounter := 0
	sharpCounter := 0
	for i, str := range strArr {
		if str == "#" {
			sharpCounter++
		} else {
			numCounter++
		}
		if sharpCounter > numCounter+1 || sharpCounter == numCounter+1 && i != len(strArr)-1 {
			return false
		}
	}
	return true
}
