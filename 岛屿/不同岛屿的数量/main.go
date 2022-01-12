package main

// https://leetcode-cn.com/problems/number-of-distinct-islands/
func main() {

}

// 1 岛屿
// 0 海水
func numDistinctIslands(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	islands := make(map[string]bool) // 去重
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				path := make([]string, 0)
				dfs(grid, i, j, m, n, &path, "666") // 初始化dir写什么都可以
				s := constructString(path)
				if _, ok := islands[s]; !ok {
					islands[s] = true
				}
			}
		}
	}
	return len(islands)
}

func constructString(str []string) string {
	res := ""
	for _, s := range str {
		res = res + s
	}
	return res
}

func dfs(grid [][]int, i, j, m, n int, path *[]string, dir string) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	*path = append(*path, dir)
	*path = append(*path, ",")

	dfs(grid, i-1, j, m, n, path, "u")
	dfs(grid, i+1, j, m, n, path, "d")
	dfs(grid, i, j-1, m, n, path, "l")
	dfs(grid, i, j+1, m, n, path, "r")

	*path = append(*path, "-"+dir)
	*path = append(*path, ",")
}
