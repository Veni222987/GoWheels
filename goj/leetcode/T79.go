package leetcode

// ? 为啥这个解法会超时呢
// 单词搜索
func exist(board [][]byte, word string) bool {
	type coord struct {
		i int
		j int
	}
	deepCopy := func(originalMap map[coord]bool) map[coord]bool {
		copiedMap := make(map[coord]bool)
		for key, value := range originalMap {
			copiedMap[key] = value
		}
		return copiedMap
	}

	var backtrack func(board [][]byte, c coord, word string, index int, visited map[coord]bool) bool
	backtrack = func(board [][]byte, c coord, word string, index int, visited map[coord]bool) bool {
		m, n := len(board), len(board[0])
		_, ok := visited[c]
		if c.i < 0 || c.j < 0 || c.i >= m || c.j >= n ||
			ok || board[c.i][c.j] != word[index] {
			return false
		}
		visited[c] = true
		if index == len(word)-1 {
			return true
		}
		return backtrack(board, coord{c.i - 1, c.j}, word, index+1, deepCopy(visited)) ||
			backtrack(board, coord{c.i + 1, c.j}, word, index+1, deepCopy(visited)) ||
			backtrack(board, coord{c.i, c.j - 1}, word, index+1, deepCopy(visited)) ||
			backtrack(board, coord{c.i, c.j + 1}, word, index+1, deepCopy(visited))
	}

	for i := range board {
		for j := range board[i] {
			if backtrack(board, coord{i, j}, word, 0, make(map[coord]bool)) {
				return true
			}
		}
	}
	return false
}
