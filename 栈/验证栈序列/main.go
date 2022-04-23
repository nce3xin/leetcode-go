package main

// 946
func validateStackSequences(pushed []int, popped []int) bool {
	// 利用栈进行模拟
	s := make([]int, 0)
	//i是pushed的索引，j是popped的索引
	i := 0
	j := 0
	for ; i < len(pushed); i++ {
		//只要i还没到头，就先push一个先
		s = append(s, pushed[i])
		//当栈顶元素等于popped[j], 出栈，j后移一位
		for len(s) > 0 && s[len(s)-1] == popped[j] {
			s = s[:len(s)-1]
			j++
		}
	}
	//若栈为空，即表示符合要求返回true
	if len(s) == 0 {
		return true
	}
	return false
}
