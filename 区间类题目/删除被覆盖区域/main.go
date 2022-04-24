package main

import "sort"

// 1288

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)
	// 按照起点升序排列，起点相同时降序排列
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	// 记录合并区间的起点和终点
	left := intervals[0][0]
	right := intervals[0][1]
	numOverlapped := 0
	for i := 1; i < n; i++ {
		curLeft := intervals[i][0]
		curRight := intervals[i][1]
		// 情况1：找到覆盖区间
		if curLeft >= left && curRight <= right {
			numOverlapped++
		}
		// 情况2：找到相交区间，合并
		if curLeft >= left && curRight >= right {
			right = curRight
		}
		// 情况3：完全不相交，更新起点和终点
		if curLeft > right {
			left = curLeft
			right = curRight
		}
	}
	return n - numOverlapped
}
