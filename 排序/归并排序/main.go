package main

// 912
func sortArray(nums []int) []int {
	n := len(nums)
	// 排序整个数组（原地修改）
	mergeSort(nums, 0, n-1)
	return nums
}

// 闭区间 [lo, hi]
// 归并排序其实就是二叉树的后续遍历
// 定义：将子数组 nums[lo..hi] 进行排序
func mergeSort(nums []int, lo, hi int) {
	if lo >= hi {
		// 单个元素不用排序
		return
	}
	// 这样写是为了防止溢出，效果等同于 (hi + lo) / 2
	mid := lo + (hi-lo)/2
	// 先对左半部分数组 nums[lo..mid] 排序
	mergeSort(nums, lo, mid)
	// 再对右半部分数组 nums[mid+1..hi] 排序
	mergeSort(nums, mid+1, hi)
	// 将两部分有序数组合并成一个有序数组
	merge(nums, lo, mid, hi)
}

// 将 nums[lo..mid] 和 nums[mid+1..hi] 这两个有序数组合并成一个有序数组
func merge(nums []int, lo, mid, hi int) {
	// 指向两部分起点
	i := lo
	j := mid + 1
	// 由于要原地排序，所以需要一个辅助数组
	res := make([]int, 0)
	// 两部分都没有遍历完
	for i <= mid && j <= hi {
		if nums[i] < nums[j] {
			res = append(res, nums[i])
			i++
		} else {
			res = append(res, nums[j])
			j++
		}
	}
	// 第一部分还有数据，第二部分已经没有数据了
	for i <= mid {
		res = append(res, nums[i])
		i++
	}
	// 第一部分没有数据了，第二部分还有
	for j <= hi {
		res = append(res, nums[j])
		j++
	}
	copy(nums[lo:hi+1], res)
}
