package main

import "fmt"

func main() {
	nums := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	res := sortArray(nums)
	fmt.Printf("res: %v\n", res)
}

// 912
func sortArray(nums []int) []int {
	n := len(nums)
	quickSort(nums, 0, n-1)
	return nums
}

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIndex := partition(nums, lo, hi)
	quickSort(nums, lo, pivotIndex-1)
	quickSort(nums, pivotIndex+1, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	less := lo - 1
	more := hi
	cur := lo
	for cur < more {
		if nums[cur] < pivot {
			nums[cur], nums[less+1] = nums[less+1], nums[cur]
			less++
			cur++
		} else if nums[cur] > pivot {
			// 这儿是交换nums[more-1], 不是nums[hi-1]！
			nums[cur], nums[more-1] = nums[more-1], nums[cur]
			more--
		} else {
			cur++
		}
	}
	// 交换pivot和大于区域的第一个数
	nums[hi], nums[more] = nums[more], nums[hi]
	// 返回交换后，pivot所在位置的下标
	return more
}
