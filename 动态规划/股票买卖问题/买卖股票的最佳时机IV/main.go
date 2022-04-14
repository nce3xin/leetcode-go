package main

import "math"

// 188

func maxProfit(k int, prices []int) int {
	// dp[i][j][k]: 前i天，交易j次，最终持有状态是k（1表示持有，0表示未持有）。能获得的最大利润
	n := len(prices)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// base case
	// dp[0][...][0]: 前0天，最终未持有，利润为0
	// dp[0][...][1]: 前0天，最终持有，不可能。题目求最大值，这里给个最小值，方便求最大值
	// dp[...][0][0]: 交易0次，最终未持有，利润为0
	// dp[...][0][1]: 交易0次，最终持有，不可能。题目求最大值，这里给一个最小值，方便求最大值
	for j := 0; j <= k; j++ {
		dp[0][j][1] = math.MinInt32
	}
	for i := 0; i <= n; i++ {
		dp[i][0][1] = math.MinInt32
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(
				dp[i-1][j][0],
				dp[i-1][j][1]+prices[i-1],
			)
			dp[i][j][1] = max(
				dp[i-1][j][1],
				dp[i-1][j-1][0]-prices[i-1],
			)
		}
	}
	return dp[n][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
