package main

// 875

// 二分法
// 第一步：确定x, f(x), target。并写出f(x)的代码
// 第二步，找到x的取值范围，作为二分搜索的搜索区间。初始化 left 和 right 变量
func minEatingSpeed(piles []int, h int) int {
	// 二分搜索区间：[left, right)
	left := 1
	// 这里不能这么写：10^9，因为在go语言中，^表示异或符号
	right := 1000000000 + 1
	for left < right {
		mid := left + (right-left)/2
		if f(piles, mid) == h {
			// 求吃完所有香蕉的最小速度，因此是求左边界
			right = mid
		} else if f(piles, mid) < h {
			right = mid
		} else if f(piles, mid) > h {
			left = mid + 1
		}
	}
	return left
}

// x；speed. 吃香蕉的速度
// f(x): 吃完所有香蕉的用时
// f(x) 是 x 的单调递减函数
// target 也很明显。在用时小于等于 h 的限制条件下，求 x 的最小值。所以 target 就是 h。
func f(piles []int, speed int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / speed
		if piles[i]%speed > 0 {
			hours++
		}
	}
	return hours
}
