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
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		num := candidates[i]
		*track = append(*track, num)
		// 和v1版本的区别在这儿，前后不单独处理sum加减，直接在调用backtrack传参的地方计算，也能ac
		backtrack(candidates, i, target, sum+num, track, res)
		*track = (*track)[:len(*track)-1]
	}
}
