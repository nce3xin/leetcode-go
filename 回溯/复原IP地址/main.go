package main

import "strconv"

// 93

// 回溯
func restoreIpAddresses(s string) []string {
	track := ""
	res := make([]string, 0)
	cnt := 0
	n := len(s)
	backtrack(&s, &track, &res, &cnt, n)
	return res
}

// 没有用传参start的方式，而是选择不断缩短s，这样写起来逻辑简单一点
// 参数 length：记录原始s的长度。之所以要记录，是因为s在不断收缩。
func backtrack(s *string, track *string, res *[]string, cnt *int, length int) {
	// 加1是因为这时track末尾有'.'
	if *cnt == 4 && len(*track) == length+3+1 {
		// 去除track末尾的'.'
		t := (*track)[:len(*track)-1]
		*res = append(*res, t)
		return
	}
	// 心里画出一棵多叉树（递归树）
	// 做选择，当前s的每个字符后面都可以尝试加分隔符'.'
	for i := 0; i < len(*s); i++ {
		cur := (*s)[:i+1]
		if !isValid(cur) {
			continue
		}
		*cnt++
		oldLength := len(*track)
		*track += cur + "."
		// 这里必须新开一个变量subs，递归时传递&subs。不能*s = (*s)[i+1:]，然后传递s，否则最终结果是空
		subs := (*s)[i+1:]
		backtrack(&subs, track, res, cnt, length)
		*track = (*track)[:oldLength]
		*cnt--
	}
}

func isValid(s string) bool {
	if s == "0" {
		return true
	}
	if len(s) > 3 {
		return false
	}
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num < 0 || num > 255 {
		return false
	}
	return true
}
