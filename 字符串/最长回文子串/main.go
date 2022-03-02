package main

// 5

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := findPalindrome(s, i, i)
		s2 := findPalindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

// 找到以l,r为中心的最长回文子串
func findPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return s[l+1 : r] // 左闭右开
}
