package main

// 198

func rob(nums []int) int {
	// dp[i]: 盗窃前i个房屋，能偷到的最大金额
	n := len(nums)
	dp := make([]int, n+1)
	// base case
	// dp[0]: 没有盗窃，金额为0
	// dp[1]: 盗窃第1个物资，能偷到的金额为nums[0]
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
