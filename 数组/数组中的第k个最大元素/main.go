package main

// 215
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 找第k大元素，就是找升序排序后，下标为n-k的元素
	targetIndex := n - k
	lo := 0
	hi := n - 1
	for lo <= hi {
		// 每partition一次，就会排好一个数组元素，就是pivot。这里返回排好序之后的pivot的下标
		p := partition(nums, lo, hi)
		if p < targetIndex {
			// 是 lo = p + 1，不是 lo = targetIndex + 1
			lo = p + 1
		} else if p > targetIndex {
			hi = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	// less是小于区域的右边界（闭区间）
	less := lo - 1
	// more是大于区域的左边界（闭区间）
	more := hi
	cur := lo
	// 这里是for，不是if！每次都在这种问题上犯蠢
	for cur < more {
		if nums[cur] < pivot {
			nums[cur], nums[less+1] = nums[less+1], nums[cur]
			less++
			cur++
		} else if nums[cur] > pivot {
			nums[cur], nums[more-1] = nums[more-1], nums[cur]
			// 这里是more--，不是more++!
			more--
		} else {
			cur++
		}
	}
	// 交换pivot和大于区域的第一个数
	nums[hi], nums[more] = nums[more], nums[hi]
	return more
}
