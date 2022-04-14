package main

import "math"

// 309
func maxProfit(prices []int) int {
	// dp[i][j]: 前i天各种操作，最终持有状态是j（1表示持有，0表示未持有）。能获得的最大利润。
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	// base case
	// dp[0][0]: 前i天，最终未持有，利润为0
	// dp[0][1]: 前i天，最终持有，不可能。题目求最大值，这里给一个最小值，方便求最大值
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		dp[i][0] = max(
			dp[i-1][0],
			dp[i-1][1]+prices[i-1],
		)
		if i-2 >= 0 {
			dp[i][1] = max(
				dp[i-1][1],
				dp[i-2][0]-prices[i-1],
			)
		} else {
			// 注意，这里也是求max, 而不是 dp[i][1] = dp[i-1][1]
			dp[i][1] = max(dp[i-1][1], -prices[i-1])
		}
	}
	return dp[n][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
