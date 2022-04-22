package main

// 848
// 后缀和数组
func shiftingLetters(s string, shifts []int) string {
	n := len(s)
	bytes := make([]byte, 0)
	// suffixSum[i]: s[i]需要移动的总次数
	suffixSum := make([]int, n)
	suffixSum[n-1] = shifts[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + shifts[i]
	}
	for i := 0; i < n; i++ {
		moveCnt := suffixSum[i]
		c := s[i]
		moveCnt = moveCnt % 26
		newC := byte(int('a') + (int(c)-int('a')+moveCnt)%26)
		bytes = append(bytes, newC)
	}
	return string(bytes)
}
