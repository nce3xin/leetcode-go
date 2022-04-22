package main

// 二分法
// 二分，参数里不用lo, hi这些。那是快排、归并才用的
func search(nums []int, target int) int {
	n := len(nums)
	// 搜索区间，闭区间[left,right]
	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 必须小于等于
		if nums[left] <= nums[mid] { // 左区间有序
			// target < nums[mid]，不能是 target <= nums[mid - 1], 因为 mid-1 可能会越界
			if target >= nums[left] && target < nums[mid] { // target在左区间
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右区间有序
			// target > nums[mid]，不能是 target >= nums[mid + 1]，因为 mid + 1 可能会越界
			if target > nums[mid] && target <= nums[right] { // target在右区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
