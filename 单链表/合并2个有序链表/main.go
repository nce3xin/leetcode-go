package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	// 合并后 p1 和 p2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
