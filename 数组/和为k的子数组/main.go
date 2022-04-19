package main

// 560

// 这个解法，时间复杂度O(N^2)，空间复杂度O(N)，并不是最优解
func subarraySum(nums []int, k int) int {
	n := len(nums)
	prefixSum := preSum(nums, n)
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if sumRange(prefixSum, i, j) == k {
				res++
			}
		}
	}
	return res
}

// 前缀和数组, res[i]表示前i个元素的累加和，即nums[0...i-1]的累加和
func preSum(nums []int, n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i-1] + nums[i-1]
	}
	return res
}

// 求闭区间[left, right]的累加和
func sumRange(preSum []int, left, right int) int {
	return preSum[right+1] - preSum[left]
}
