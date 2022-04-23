package main

import "fmt"

// 143

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{1, 2, 3, 4, 5}
	head := buildList(nums)
	printList(head)
	reorderList(head)
	printList(head)
}

func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%v ", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func buildList(nums []int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummyHead
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}

func reorderList(head *ListNode) {
	mid := findMiddle(head)
	newHead := reverseList(mid.Next)
	// mid.Next = nil不能在reverseList前面，否则参数mid.Next就永远为空了
	mid.Next = nil
	mergeList(head, newHead)
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 链表长度为偶数
	if fast == nil {
		return prev
	}
	return slow
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

func mergeList(head1, head2 *ListNode) {
	p1 := head1
	p2 := head2
	for p1 != nil && p2 != nil {
		// 先保存p1.Next和p2.Next，因为后面要修改这两个值
		p1Next := p1.Next
		p2Next := p2.Next
		p1.Next = p2
		p1 = p1Next
		p2.Next = p1
		p2 = p2Next
	}
}
