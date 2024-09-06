package leetcode

type coord struct {
	i int
	j int
}

// leetcode994 mid 腐烂的橘子
func orangesRotting(grid [][]int) int {
	good := 0
	queue := []coord{}
	// 找出所有腐烂的橘子
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				good++
			}
			if grid[i][j] == 2 {
				queue = append(queue, coord{i: i, j: j})
			}
		}
	}
	minute := 0
	// 层序遍历腐烂
	for len(queue) > 0 && good > 0 {
		minute++
		l := len(queue)
		for i := 0; i < l; i++ {
			c := queue[i]
			// 上下左右
			for _, d := range []coord{
				{i: c.i - 1, j: c.j},
				{i: c.i + 1, j: c.j},
				{i: c.i, j: c.j - 1},
				{i: c.i, j: c.j + 1},
			} {
				if d.i >= 0 && d.j >= 0 && d.i < len(grid) && d.j < len(grid[d.i]) && grid[d.i][d.j] == 1 {
					grid[d.i][d.j] = 2
					good--
					queue = append(queue, d)
				}
			}
		}
		queue = queue[l:]
	}
	if good > 0 {
		return -1
	}
	return minute
}
