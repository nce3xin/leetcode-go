package main

// 25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [p, q) 包含 k 个待反转元素
	p := head
	q := head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if q == nil {
			return head
		}
		q = q.Next
	}
	// 反转前 k 个元素
	newHead := reverse(p, q)
	// 递归反转后续链表并连接起来
	p.Next = reverseKGroup(q, k)
	return newHead
}

// 反转链表左闭右开区间 [left, right) 的值
func reverse(left, right *ListNode) *ListNode {
	var prev *ListNode
	cur := left
	for cur != right {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 返回反转后的头结点
	return prev
}
