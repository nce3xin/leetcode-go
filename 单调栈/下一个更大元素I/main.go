package main

import "fmt"

// 496

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Printf("res: %v\n", res)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	// 记录nums2中元素值到下标的映射
	m := make(map[int]int)
	for i := 0; i < n2; i++ {
		num := nums2[i]
		m[num] = i
	}
	next := nextGreater(nums2)
	// fmt.Printf("next: %v\n", next)
	res := make([]int, 0)
	for i := 0; i < n1; i++ {
		num := nums1[i]
		index := m[num]
		nextGreaterVal := next[index]
		res = append(res, nextGreaterVal)
	}
	return res
}

func nextGreater(nums []int) []int {
	n := len(nums)
	// 注意，这里res长度要初始化为 n
	// 因为下面赋值是给 res[i]，而不是 res = append(res, ...), 因为是倒序的
	res := make([]int, n)
	s := make([]int, 0)
	// 倒着往栈里放
	for i := n - 1; i >= 0; i-- {
		// 判定个子高矮
		num := nums[i]
		for len(s) != 0 && s[len(s)-1] <= num {
			// 矮个起开，反正也被挡着了。。。
			s = s[:len(s)-1]
		}
		// nums[i] 身后的 next great number
		if len(s) == 0 {
			res[i] = -1
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, num)
	}
	return res
}
