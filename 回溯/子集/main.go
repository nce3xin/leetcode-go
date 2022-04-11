package main

// 78

func subsets(nums []int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(nums, 0, len(nums), &track, &res)
	return res
}

func backtrack(nums []int, start int, n int, track *[]int, res *[][]int) {
	// 因为回溯树每一层都是解，所以这里直接添加，没有添加条件了
	t := make([]int, len(*track))
	copy(t, *track)
	*res = append(*res, t)
	// 遍历选择列表
	// 每次只能选择start后面的元素，因为一旦选择前面的元素，会出现重复结果
	for i := start; i < n; i++ {
		num := nums[i]
		*track = append(*track, num)
		// 这里backtrack的参数，是i+1，不是start+1，别搞错了
		backtrack(nums, i+1, n, track, res)
		*track = (*track)[:len(*track)-1]
	}
}
