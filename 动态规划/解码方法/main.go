package main

// 91
func numDecodings(s string) int {
	n := len(s)
	// dp[i]: 表示字符串前i个字符的解码方法总数
	dp := make([]int, n+1)
	// base case
	// 前0个字符，解码方法数为1
	dp[0] = 1 // 注意这里赋值是1不是0，空字符串可以有 1 种解码方法，解码出一个空字符串
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= n; i++ {
		cur := s[i-1]
		prev := s[i-2]
		if cur >= '1' && cur <= '9' {
			dp[i] += dp[i-1]
		}
		// 注意这里的运算优先级
		if prev == '1' || (prev == '2' && (cur >= '0' && cur <= '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
