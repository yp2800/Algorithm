package leetcode

// 尝试优先搜索
func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}

	nc := len(grid[0])
	nums_islands := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				nums_islands++
				dfs(grid, r, c)
			}
		}
	}
	return nums_islands
}

func dfs(grid [][]byte, r, c int) {
	nr := len(grid)
	nc := len(grid[0])

	grid[r][c] = '0'
	if r-1 >= 0 && grid[r-1][c] == '1' {
		dfs(grid, r-1, c)
	}
	if r+1 < nr && grid[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == '1' {
		dfs(grid, r, c-1)
	}
	if c+1 < nc && grid[r][c+1] == '1' {
		dfs(grid, r, c+1)
	}
}

// 广度优先搜索
func numIslands2(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}

	queue := make([]position, 0)
	nc := len(grid[0])
	nums_islands := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				nums_islands++
				grid[r][c] = '0'
				queue = append(queue, position{r, c})
				for len(queue) != 0 {
					row := queue[0].rowIndex
					col := queue[0].colIndex
					queue = queue[1:]
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, position{row - 1, col})
						grid[row-1][col] = '0'
					}
					if row+1 < nr && grid[row+1][col] == '1' {
						queue = append(queue, position{row + 1, col})
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, position{row, col - 1})
						grid[row][col-1] = '0'
					}
					if col+1 < nc && grid[row][col+1] == '1' {
						queue = append(queue, position{row, col + 1})
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}
	return nums_islands
}

type position struct {
	rowIndex, colIndex int
}
