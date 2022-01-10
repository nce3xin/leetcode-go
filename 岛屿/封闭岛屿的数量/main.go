package main

// 1254 题

func main() {

}

func closedIsland(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	floodEdgeIsland(grid, m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j, m, n)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	// 0是陆地，1是海水
	if grid[i][j] == 1 {
		return
	}
	// 淹了
	grid[i][j] = 1
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}

func floodEdgeIsland(grid [][]int, m, n int) {
	// 把上方和下方淹了以及连着它们的岛屿淹了
	// 用dfs来淹，不能只淹一块，要把连着的都淹掉，所以要用dfs
	for j := 0; j < n; j++ {
		// 最上边淹了
		dfs(grid, 0, j, m, n)
		// 最下边淹了
		dfs(grid, m-1, j, m, n)
	}
	// 把左边和右边淹了
	for i := 0; i < m; i++ {
		// 最左边淹了
		dfs(grid, i, 0, m, n)
		// 最右边淹了
		dfs(grid, i, n-1, m, n)
	}
}
