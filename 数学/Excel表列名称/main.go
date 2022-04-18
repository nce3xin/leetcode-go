package main

import "fmt"

// 168

func main() {
	columnNumber := 27
	res := convertToTitle(columnNumber)
	fmt.Printf("res: %v\n", res)
}

func convertToTitle(columnNumber int) string {
	// 用栈保存余数
	s := make([]int, 0)
	for columnNumber != 0 {
		// for循环内，columnNumber先减1
		columnNumber -= 1
		// 余数
		remainder := columnNumber % 26
		columnNumber = columnNumber / 26
		s = append(s, remainder)
	}
	res := ""
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		res += string(num2char(cur))
	}
	return res
}

func num2char(num int) byte {
	res := byte('A' + num)
	return res
}
