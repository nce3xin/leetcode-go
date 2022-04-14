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
	p, q := partition(nums, lo, hi)
	quickSort(nums, lo, p)
	quickSort(nums, q, hi)
}

func partition(nums []int, lo, hi int) (int, int) {
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
			nums[cur], nums[more-1] = nums[more-1], nums[cur]
			more--
		} else {
			cur++
		}
	}
	// 交换pivot和大于区域的第一个数
	nums[hi], nums[more] = nums[more], nums[hi]
	// 返回小于区域的右边界（包含）和大于区域的左边界（包含）
	return less, more + 1
}
