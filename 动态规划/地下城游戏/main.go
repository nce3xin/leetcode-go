package main

import "math"

// 174

func calculateMinimumHP(dungeon [][]int) int {
	// dp[i][j]: 从grid[i][j]走到右下角，所需的最小血量
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// base case
	// dp[m-1][n-1]: 如果grid[m-1][n-1]大于等于0，说明所需的最小血量是1，否则为 -grid[m-1][n-1] + 1
	initDp(&dp, m, n, dungeon)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			// 最后一行
			if i == m-1 {
				dp[i][j] = dp[i][j+1] - dungeon[i][j]
				dp[i][j] = process(dp[i][j])
				continue
			}
			if j == n-1 {
				dp[i][j] = dp[i+1][j] - dungeon[i][j]
				dp[i][j] = process(dp[i][j])
				continue
			}
			dp[i][j] = min(
				dp[i][j+1], // 右方
				dp[i+1][j], // 下方
			) - dungeon[i][j]
			dp[i][j] = process(dp[i][j])
		}
	}
	return process(dp[0][0])
}

func initDp(d *[][]int, m, n int, dungeon [][]int) {
	dp := *d
	if dungeon[m-1][n-1] >= 0 {
		dp[m-1][n-1] = 1
	} else {
		dp[m-1][n-1] = -dungeon[m-1][n-1] + 1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func process(minHp int) int {
	if minHp <= 0 {
		return 1
	}
	return minHp
}
