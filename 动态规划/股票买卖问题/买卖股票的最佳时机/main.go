package main

import "math"

//121

// 注意，此题只能买卖一次！
func maxProfit(prices []int) int {
	// dp[i][j]: 前i天各种操作，最后持有状态为j（1表示持有，0表示未持有），获得的最大利润
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	// base case
	// dp[0][0]: 前0天，最终未持有，利润为0
	// dp[0][1]: 前0天，最终持有，不可能，因为求最大值，所以这里给一个最小值，方便求最大值
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		dp[i][0] = max(
			dp[i-1][1]+prices[i-1],
			dp[i-1][0],
		)
		dp[i][1] = max(
			-prices[i-1], // 之所以不是dp[i-1][0] - prices[i-1], 是因为只允许一次交易。
			dp[i-1][1],
		)
	}
	return dp[n][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
