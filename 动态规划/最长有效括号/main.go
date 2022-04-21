package main

// 32
func longestValidParentheses(s string) int {
	res := 0
	n := len(s)
	// dp[i]: 以nums[i]结尾的最长有效括号的长度
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			dp[i] = 0
		}
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else { // s[i-1] == ')'
				if i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
