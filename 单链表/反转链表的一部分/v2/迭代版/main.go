package main

// 头插法

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头插法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 虚拟头节点，方便后序操作。为了简化此情况：当left为1时，也就是从头开始反转
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummyHead.Next = head
	// guard指针，指向要反转的第一个节点的前一个节点。后续guard指针保持不变
	guard := dummyHead
	// cur指针，指向要反转的第一个节点
	cur := dummyHead
	// 移动guard和cur到指定位置
	for i := 0; i < left-1; i++ {
		guard = guard.Next
	}
	cur = guard.Next
	// 开始头插法
	for i := 0; i < right-left; i++ {
		removed := cur.Next
		cur.Next = cur.Next.Next
		// 下面这两句顺序不能变
		removed.Next = guard.Next
		guard.Next = removed
	}
	return dummyHead.Next
}
