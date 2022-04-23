package main

import (
	"fmt"
	"math"
)

func main() {
	s := "9223372036854775808"
	res := myAtoi(s)
	fmt.Printf("res: %v\n", res)
}

// 8
func myAtoi(s string) int {
	n := len(s)
	// 去除前导空格
	i := 0
	for ; i < n; i++ {
		if s[i] != ' ' {
			break
		}
	}
	s = s[i:]
	// 默认符号是正
	sign := 1
	res := 0
	for index, c := range s {
		// 转整数
		if c >= '0' && c <= '9' {
			res = 10*res + int(c-'0')
		} else if c == '-' && index == 0 {
			sign = -1
		} else if c == '+' && index == 0 {
			sign = 1
		} else {
			break
		}
		// 边界处理
		// 注意这两句if的位置，必须要再for循环里面
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -res < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * res
}
