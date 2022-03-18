package main

import "fmt"

// 300

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Printf("res: %v\n", res)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 注意这里必须要有判断
			// 为了保障元素单调递增，只能从i前面且末尾元素比nums[i]小的状态转移过来
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := dp[0]
	for i := 1; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
