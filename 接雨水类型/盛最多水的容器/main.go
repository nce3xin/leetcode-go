package main

// 11
func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	res := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		// 每次都移动短板
		// 因为移动短板，水槽的短板可能变大，因此下个水槽的面积可能增大
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
