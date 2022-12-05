package v2

// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	stack := make([]byte, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		c := s[i]
		if isLeft(c) {
			stack = append(stack, c)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == LeftOf(c) {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func isLeft(c byte) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	}
	return false
}

func LeftOf(c byte) byte {
	var res byte
	if c == ')' {
		res = '('
	} else if c == ']' {
		res = '['
	} else if c == '}' {
		res = '{'
	}
	return res
}
