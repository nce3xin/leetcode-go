package main

//695
func main() {

}

// 1 陆地
// 0 海水
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(grid, i, j, m, n))
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j, m, n int) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	// 已经是海水了
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs(grid, i-1, j, m, n) +
		dfs(grid, i+1, j, m, n) +
		dfs(grid, i, j-1, m, n) +
		dfs(grid, i, j+1, m, n) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
