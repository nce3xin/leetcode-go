package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != nil || p2 != nil {
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
			continue
		}
		if p2 == nil {
			p2 = headA
			continue
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}
