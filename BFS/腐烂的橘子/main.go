package main

type Location struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	q := make([]Location, 0)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	starts := findAllStart(grid, m, n, &q)
	for i := 0; i < len(starts); i++ {
		loc := starts[i]
		visited[loc.x][loc.y] = true
		q = append(q, loc)
	}
	numFresh := freshOrangeNum(grid, m, n)
	if numFresh == 0 {
		return 0
	}
	res := 0
	for len(q) > 0 {
		if numFresh == 0 {
			return res
		}
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			x := cur.x
			y := cur.y
			// 上方
			if x-1 >= 0 && !visited[x-1][y] && grid[x-1][y] == 1 {
				q = append(q, Location{x: x - 1, y: y})
				visited[x-1][y] = true
				numFresh--
			}
			// 下方
			if x+1 < m && !visited[x+1][y] && grid[x+1][y] == 1 {
				q = append(q, Location{x: x + 1, y: y})
				visited[x+1][y] = true
				numFresh--
			}
			// 左方
			if y-1 >= 0 && !visited[x][y-1] && grid[x][y-1] == 1 {
				q = append(q, Location{x: x, y: y - 1})
				visited[x][y-1] = true
				numFresh--
			}
			// 右方
			if y+1 < n && !visited[x][y+1] && grid[x][y+1] == 1 {
				q = append(q, Location{x: x, y: y + 1})
				visited[x][y+1] = true
				numFresh--
			}
		}
		res++
	}
	if numFresh > 0 {
		return -1
	}
	return res
}

func freshOrangeNum(grid [][]int, m, n int) int {
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 发现新鲜橘子
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func findAllStart(grid [][]int, m, n int, q *[]Location) []Location {
	res := make([]Location, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 找到腐烂的橘子
			if grid[i][j] == 2 {
				res = append(res, Location{x: i, y: j})
			}
		}
	}
	return res
}
