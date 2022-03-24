package main

// 213

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(
		robRange(nums, 1, n-1),
		robRange(nums, 0, n-2),
	)
}

// 闭区间[i,j]
func robRange(nums []int, i, j int) int {
	// dp[i]: 盗窃前i个房屋，能偷到的最大金额
	nums = nums[i : j+1]
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
