package main

// 206

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 迭代解法
func reverseListIterative(head *ListNode) *ListNode {
	// 用 prev 和 next 保存前一个节点和后一个节点
	var prev, next *ListNode
	cur := head
	for cur != nil {
		// 先保存后一个节点
		next = cur.Next
		// 修改当前节点cur的next指针为prev
		cur.Next = prev
		// 修改prev为当前节点
		prev = cur
		// cur指针向后移动一个节点
		// 注意这里不能是 cur = cur.Next，因为cur.Next已经被修改为前一个节点
		cur = next
	}
	return prev
}
