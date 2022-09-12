package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 92
// 迭代
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 定义一个dummyHead, 方便处理
	dummyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummyHead.Next = head
	// guard 第一个要反转节点的前一个节点
	// 初始化指针
	guard := dummyHead
	cur := dummyHead.Next
	// 将指针移到相应的位置
	for i := 0; i < left-1; i++ {
		guard = guard.Next
		cur = cur.Next
	}
	// 头插法插入节点
	// 这个for循环里，没有显式的cur向后移动一位，是因为头插法！cur前面不断的在插入，相当于cur后移了！
	for i := 0; i < right-left; i++ {
		removed := cur.Next
		cur.Next = cur.Next.Next
		// 不能是 removed.Next = cur, 否则解答错误
		// 第一次没问题，当cur往后走了很远的时候，cur就不是guard后面一个了。因为guard是固定的，cur是不断移动的
		// 下面这两句顺序不能变
		removed.Next = guard.Next
		guard.Next = removed
	}
	return dummyHead.Next
}
