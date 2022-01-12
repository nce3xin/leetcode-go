package main

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	p2 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
