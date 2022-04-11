package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Printf("res: %v\n", res)
}

func permuteUnique(nums []int) [][]int {
	// 区别1：首先对数组排序，使相同元素相邻
	sort.Ints(nums)
	n := len(nums)
	track := make([]int, 0)
	res := make([][]int, 0)
	visited := make([]bool, n)
	backtrack(nums, &track, &res, &visited)
	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, visited *[]bool) {
	if len(*track) == len(nums) {
		t := make([]int, len(*track))
		copy(t, *track)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] {
			continue
		}
		// 区别2：新增一个剪枝条件。当前一个相同元素没有被使用过时，跳过。
		// 只有当前一个相同元素被使用过时，才继续递归选择。
		// 这样就保证了相同元素在排序中的相对位置固定
		if i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1] {
			continue
		}
		(*visited)[i] = true
		*track = append(*track, nums[i])
		backtrack(nums, track, res, visited)
		(*visited)[i] = false
		*track = (*track)[:len(*track)-1]
	}
}
