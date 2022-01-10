package main

//1905
func main() {

}

// 1 陆地
// 0 海水
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	m := len(grid1)
	n := len(grid1[0])
	floodIslandNotInGrid1(grid1, grid2, m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(grid2, i, j, m, n)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}

func floodIslandNotInGrid1(grid1 [][]int, grid2 [][]int, m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(grid2, i, j, m, n)
			}
		}
	}
}
