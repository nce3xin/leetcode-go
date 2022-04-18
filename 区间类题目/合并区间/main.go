package main

import "sort"

// 56
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	res = append(res, intervals[0])
	n := len(intervals)
	for i := 1; i < n; i++ {
		cur := intervals[i]
		curStart := cur[0]
		curEnd := cur[1]
		prevEnd := res[len(res)-1][1]
		if curStart <= prevEnd {
			// 注意，这个地方要有一个max操作
			res[len(res)-1][1] = max(curEnd, prevEnd)
		} else {
			res = append(res, cur)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
