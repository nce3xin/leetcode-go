package main

// 34
func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	n := len(nums)
	// 左闭右开区间 [left, right)
	left := 0
	right := n
	// 循环条件为什么是 left < right，不是 left <= right？
	// 因为退出条件是 left == right, 此时搜索区间 [left, left) 为空，所以可以正确终止。
	for left < right {
		// 防止整数溢出
		mid := left + (right-left)/2
		if target > nums[mid] {
			// 在 [mid+1, right) 区间继续搜索
			left = mid + 1
		} else if target < nums[mid] {
			// 在 [left, mid) 区间继续搜索
			right = mid
		} else if target == nums[mid] {
			// 不要立即返回，而是缩小「搜索区间」的上界 right，
			// 在区间 [left, mid) 中继续搜索，即不断向左收缩，达到锁定左侧边界的目的。
			right = mid
		}
	}
	// 为什么返回 left 而不是 right？
	// 都是一样的，因为 while 终止的条件是 left == right
	// 正确返回-1，需要两个if条件
	// target 比所有数都大
	if left == n {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findLast(nums []int, target int) int {
	n := len(nums)
	// 左闭右开区间 [left, right)
	left := 0
	right := n
	// 循环条件为什么是 left < right，不是 left <= right？
	// 因为退出条件是 left == right, 此时搜索区间 [left, left) 为空，所以可以正确终止。
	for left < right {
		// 注意整数溢出
		mid := left + (right-left)/2
		if target > nums[mid] {
			// 在 [mid+1, right) 区间继续搜索
			left = mid + 1
		} else if target < nums[mid] {
			// 在 [left, mid) 区间继续搜索
			right = mid
		} else if target == nums[mid] {
			// 不要立即返回，而是缩小「搜索区间」的下界 left，
			// 在区间 [mid + 1, right) 中继续搜索，即不断向右收缩，达到锁定右侧边界的目的。
			left = mid + 1
		}
	}
	// 注意，搜索右界的时候，结果要减1
	// 也可以返回 right - 1，因为for循环的退出条件是 left == right
	// 正确返回-1，需要两个if条件
	// target 比所有数都小
	if left == 0 {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}
