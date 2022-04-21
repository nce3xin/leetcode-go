package main

// 20
func isValid(s string) bool {
	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && leftOf(c) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func leftOf(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == ']' {
		return '['
	}
	return '{'
}
