package main

import "fmt"

// 39

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Printf("res: %v\n", res)
}

func combinationSum(candidates []int, target int) [][]int {
	track := make([]int, 0)
	res := make([][]int, 0)
	backtrack(candidates, 0, target, 0, &track, &res)
	return res
}

func backtrack(candidates []int, start int, target, sum int, track *[]int, res *[][]int) {
	if sum == target {
		t := make([]int, len(*track))
		copy(t, *track)
		*res = append(*res, t)
		return
	}
	// 这个base case是需要的
	// 如果去掉，会出现回溯永远无法终止的问题，导致stack overflow
	if sum > target {
		return
	}
	// 注意，只要是组合题，i一定是需要从start开始的
	for i := start; i < len(candidates); i++ {
		num := candidates[i]
		*track = append(*track, num)
		sum += num
		// 这里调用backtrack，参数是i，不是i+1。
		// 那么反过来，如果我想让每个元素被重复使用，我只要把 i + 1 改成 i 即可：
		backtrack(candidates, i, target, sum, track, res)
		*track = (*track)[:len(*track)-1]
		sum -= num
	}
}
