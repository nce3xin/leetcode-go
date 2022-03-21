package main

// 416

func canPartition(nums []int) bool {
	s := sum(nums)
	if s%2 != 0 {
		return false
	}
	s /= 2
	n := len(nums)
	// dp[i][j]: 从前 i 个数中选择，是否存在满足和为 j 的组合。true表示存在，false表示不存在
	dp := make([][]bool, n+1)
	// base case
	// dp[0][...]: 没有数可以选择，false
	// dp[...][0]: 和为0，可以什么数都不选择，true
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1) // 默认值是false
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= s; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][s]
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}
