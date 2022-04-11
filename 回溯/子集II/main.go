package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(nums, n, 0, &track, &res)
	return res
}

func backtrack(nums []int, n int, start int, track *[]int, res *[][]int) {
	t := make([]int, len(*track))
	copy(t, *track)
	*res = append(*res, t)
	for i := start; i < n; i++ {
		num := nums[i]
		// 注意这里！剪枝条件是i > start, 而不是i > 0！
		if i > start && num == nums[i-1] {
			continue
		}
		*track = append(*track, num)
		backtrack(nums, n, i+1, track, res)
		*track = (*track)[:len(*track)-1]
	}
}
