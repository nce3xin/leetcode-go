package main

import "fmt"

// 494

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	res := findTargetSumWays(nums, target)
	fmt.Printf("res: %v\n", res)
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	res := 0
	backtrack(nums, target, &sum, 0, &res)
	return res
}

func backtrack(nums []int, target int, sum *int, i int, res *int) {
	if i == len(nums) {
		if *sum == target {
			*res += 1
		}
		return
	}
	val := nums[i]
	// 添加正号
	*sum += val
	backtrack(nums, target, sum, i+1, res)
	*sum -= val
	// 添加负号
	*sum -= val
	backtrack(nums, target, sum, i+1, res)
	*sum += val
}
