func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)

	for i := range visited{
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int) int

	dfs = func(r, c int)int {

		if r < 0 || r  >= rows || c < 0 || c >= cols {
			return 1
		} 

		if grid[r][c] == 0 {
			return 1
		}

		if visited[r][c] {
			return 0
		}

		visited[r][c] = true

		perimeter := 0
		perimeter += dfs(r - 1, c)
		perimeter += dfs(r + 1, c)
		perimeter += dfs(r, c + 1)
		perimeter += dfs(r, c - 1)

		return perimeter
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return dfs(i, j)
			}
		}
	}

	return 0
}
