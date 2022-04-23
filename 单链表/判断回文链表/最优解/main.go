package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 如果fast指针没有指向null，说明链表长度为奇数，slow还要再前进一步
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := reverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	p := head
	var prev *ListNode
	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return prev
}
