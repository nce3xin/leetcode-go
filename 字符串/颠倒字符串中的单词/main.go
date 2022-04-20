package main

import "fmt"

func main() {
	s := "  hello world  "
	res := reverseWords(s)
	fmt.Printf("res: %v\n", res)
}

// 151
func reverseWords(s string) string {
	n := len(s)
	bytes := trimSpaces(s, n)
	// 这儿有个坑，reverse函数的第二个参数，不是n-1, 而是len(bytes)-1.
	// 因为要反转的是bytes, 自然要用bytes的长度
	reverse(&bytes, 0, len(bytes)-1)
	reverseEachWord(&bytes)
	return string(bytes)
}

func reverseEachWord(bytes *[]byte) {
	n := len(*bytes)
	start := 0
	end := 0
	for start < n {
		for end < n && (*bytes)[end] != ' ' {
			end++
		}
		reverse(bytes, start, end-1)
		start = end + 1
		end = start
	}
}

func trimSpaces(s string, n int) []byte {
	p := 0
	q := n - 1
	// 删除头部的冗余空格
	for p <= q && s[p] == ' ' {
		p++
	}
	// 删除尾部的冗余空格
	for p <= q && s[q] == ' ' {
		q--
	}
	// 闭区间 [p, q] 就是去除首尾空格后的字符串
	bytes := make([]byte, 0)
	for p <= q {
		c := s[p]
		if c != ' ' {
			bytes = append(bytes, c)
		} else if bytes[len(bytes)-1] != ' ' { // 当前字符是空格，且bytes最后一个值不是空格。说明遇到第一个空格
			bytes = append(bytes, c)
		}
		p++
	}
	return bytes
}

// 反转闭区间 [left, right] 之间的元素
// 双指针原地反转
func reverse(bytes *[]byte, left, right int) {
	p := left
	q := right
	for p < q {
		(*bytes)[p], (*bytes)[q] = (*bytes)[q], (*bytes)[p]
		p++
		q--
	}
}
