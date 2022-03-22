package main

import "math"

//53

func maxSubArray(nums []int) int {
	// dp[i]: 以nums[i]元素为结尾的连续子数组的最大和
	n := len(nums)
	dp := make([]int, n)
	// base case
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	res := math.MinInt32
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
