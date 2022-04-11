package main

func combine(n int, k int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	// backtrack和backtrack2都可以，实现有一点区别。backtrack2不需要level参数
	//backtrack(n, k, 1, &track, &res, 0)
	backtrack2(n, k, 1, &track, &res)
	return res
}

// level表示层数
func backtrack(n int, k int, start int, track *[]int, res *[][]int, level int) {
	if level == k {
		t := make([]int, len(*track))
		copy(t, *track)
		*res = append(*res, t)
		return
	}
	for i := start; i <= n; i++ {
		*track = append(*track, i)
		// 注意backtrack这里是i+1，不是start+1
		backtrack(n, k, i+1, track, res, level+1)
		*track = (*track)[:len(*track)-1]
	}
}

func backtrack2(n int, k int, start int, track *[]int, res *[][]int) {
	// backtrack2 相比 backtrack，少了一个level参数，直接用track的长度是否等于k作为添加结果的条件
	if len(*track) == k {
		t := make([]int, len(*track))
		copy(t, *track)
		*res = append(*res, t)
		return
	}
	for i := start; i <= n; i++ {
		*track = append(*track, i)
		// 注意backtrack这里是i+1，不是start+1
		backtrack2(n, k, i+1, track, res)
		*track = (*track)[:len(*track)-1]
	}
}
