package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	//全部放在一起直接二分
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l != r && l != r-1 {
		mid, midNum := (l+r)>>1, 0
		if m == 1 {
			midNum = matrix[0][mid%n]
		} else {
			midNum = matrix[mid/n][mid%n]
		}

		if midNum == target {
			return true
		}
		if midNum > target {
			r = mid
		} else {
			l = mid
		}
	}
	if matrix[l/n][l%n] == target || matrix[r/n][r%n] == target {
		return true
	}
	return false
}
