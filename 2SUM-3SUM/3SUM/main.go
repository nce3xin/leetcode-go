package main

import "sort"

// 15
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		val := nums[i]
		tss := twoSum(nums, i+1, -val, n)
		if len(tss) == 0 {
			continue
		}
		for _, ts := range tss {
			res = append(res, []int{val, ts[0], ts[1]})
		}
		for i+1 < n && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start int, target int, n int) [][]int {
	p := start
	q := n - 1
	res := make([][]int, 0)
	for p < q {
		left := nums[p]
		right := nums[q]
		sum := nums[p] + nums[q]
		if sum < target {
			p++
		} else if sum > target {
			q--
		} else {
			// 注意，去重的逻辑在这里，在找到结果这里
			res = append(res, []int{nums[p], nums[q]})
			// 注意，这里一定要有p<q的判断，否则解答错误
			for p < q && left == nums[p] {
				p++
			}
			for p < q && right == nums[q] {
				q--
			}
		}
	}
	return res
}
