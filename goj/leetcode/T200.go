package leetcode

func dfs(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' || visited[i][j] {
		return
	}
	if grid[i][j] == '1' || !visited[i][j] {
		visited[i][j] = true
	}
	dfs(grid, visited, i+1, j)
	dfs(grid, visited, i-1, j)
	dfs(grid, visited, i, j+1)
	dfs(grid, visited, i, j-1)
}
func numIslands(grid [][]byte) int {
	// dfs遍历所有1
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	counter := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visited[i][j] {
				counter++
				dfs(grid, visited, i, j)
			}
		}
	}
	return counter
}
