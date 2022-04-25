package main

// 14
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	prefix := strs[0]
	for i := 1; i < n; i++ {
		cur := strs[i]
		prefix = getLcp(prefix, cur)
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func getLcp(s1, s2 string) string {
	n1 := len(s1)
	n2 := len(s2)
	length := min(n1, n2)
	i := 0
	for i < length {
		if s1[i] != s2[i] {
			break
		}
		i++
	}
	return s1[:i]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
