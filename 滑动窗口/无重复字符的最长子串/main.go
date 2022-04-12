package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	n := len(s)
	res := 0
	m := make(map[byte]int)
	for right < n {
		c := s[right]
		right++
		m[c] += 1
		for m[c] > 1 {
			m[s[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
