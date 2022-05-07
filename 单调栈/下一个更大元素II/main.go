package main

// 503
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	s := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		x := nums[i%n]
		for len(s) != 0 && s[len(s)-1] <= x {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1]
		}
		s = append(s, x)
	}
	return res
}
