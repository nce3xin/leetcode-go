package 零钱兑换

import "math"

// 322

func coinChange(coins []int, amount int) int {
	// 定义dp[amount]表示最少用多少个硬币拼出金额amount。
	dp := make([]int, amount+1)
	n := len(dp)
	for i := 0; i < n; i++ {
		// 这儿有个坑，必须是MaxInt32，不能是MaxInt64，否则解答错误。
		// 可能是因为MaxInt64加1会溢出。
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 遍历所有状态
	for i := 1; i <= amount; i++ {
		// 做选择
		// 这儿也有个坑，不能写for coin := range coins，那样coin会默认是index
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
