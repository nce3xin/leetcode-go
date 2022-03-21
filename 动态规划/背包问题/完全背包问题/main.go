package main

// 518

func change(amount int, coins []int) int {
	// dp[i][j]: 前 i 个数（可重复使用），可以凑成金额 j 的组合数
	n := len(coins)
	dp := make([][]int, n+1)
	// base case
	// dp[0][...]: 没有数可以选，0
	// dp[...][0]: 金额为0，可以什么都不选，1
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1) // 默认值为0
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}
