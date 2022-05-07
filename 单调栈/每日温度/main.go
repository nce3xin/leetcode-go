package main

// 739

type Item struct {
	Val   int
	Index int
}

// 其实就是栈中元素多保存一个索引
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	s := make([]Item, 0)
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(s) != 0 && s[len(s)-1].Val <= t {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1].Index - i
		}
		s = append(s, Item{
			Val:   t,
			Index: i,
		})
	}
	return res
}
