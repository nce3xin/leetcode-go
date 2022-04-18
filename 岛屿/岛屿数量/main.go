package main

func main() {

}

// 时间复杂度：O(M*N)，M和N分别为行数和列数
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	// 遍历 grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 每发现一个岛屿，岛屿数量加一
				res++
				// 然后使用 DFS 将岛屿淹了
				dfs(grid, i, j, m, n)
			}
		}
	}
	return res
}

// 从 (i, j) 开始，将与之相邻的陆地都变成海水
func dfs(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		// 超出索引边界
		return
	}
	if grid[i][j] == '0' {
		// 已经是海水了
		return
	}
	// 将 (i, j) 变成海水
	grid[i][j] = '0'
	// 淹没上下左右的陆地
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}
