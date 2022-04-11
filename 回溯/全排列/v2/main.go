package main

// 46

func permute(nums []int) [][]int {
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
	for i, num := range nums {
		if (*visited)[i] {
			continue
		}
		(*visited)[i] = true
		*track = append(*track, num)
		backtrack(nums, track, res, visited)
		*track = (*track)[:len(*track)-1]
		(*visited)[i] = false
	}
}
