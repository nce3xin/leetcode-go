package main

// 234

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用栈，遍历链表，入栈。然后从头一个一个比较
func isPalindrome(head *ListNode) bool {
	s := make([]*ListNode, 0)
	p := head
	for p != nil {
		s = append(s, p)
		p = p.Next
	}
	p = head
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		if p.Val != cur.Val {
			return false
		}
		p = p.Next
	}
	return true
}
