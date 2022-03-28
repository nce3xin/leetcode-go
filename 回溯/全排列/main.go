package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Printf("res: %v\n", res)
}

func permute(nums []int) [][]int {
	n := len(nums)
	track := make([]int, 0)
	res := make([][]int, 0)
	used := make([]bool, n) // default: false
	backtrack(nums, &track, &res, &used)
	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, visited *[]bool) {
	if len(nums) == len(*track) {
		// 这儿有个大坑，不能直接*res = append(*res, *track)
		// 必须要重新生成一个新的track对象
		// 因为track是一个外部引用，在遍历过程中track 中的数据会不断变化，所以装入 res 的时候应该把 track 里面的值做一次拷贝。
		t := make([]int, len(*track))
		copy(t, *track)
		*res = append(*res, t)
		return
	}
	for i, v := range nums {
		if (*visited)[i] {
			continue
		}
		*track = append(*track, v)
		(*visited)[i] = true
		backtrack(nums, track, res, visited)
		*track = (*track)[:len(*track)-1]
		(*visited)[i] = false
	}
}
