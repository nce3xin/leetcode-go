package main

import "math/rand"

// 382

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	s := Solution{head: head}
	return s
}

func (this *Solution) GetRandom() int {
	p := this.head
	i := 0
	res := 0
	for p != nil {
		i++
		// [0, i) 区间内随机选择一个整数。[0, i) 区间内的整数数量是i
		randNum := rand.Intn(i)
		// 在 [0, i) 中随机抽样，等于0的概率是 1/i
		// 1/i 的概率选择新来的值
		if randNum == 0 {
			res = p.Val
		}
		// 1 - 1/i 的概率，保持原值不变
		p = p.Next
	}
	return res
}
