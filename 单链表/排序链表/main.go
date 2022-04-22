package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	n := len(nums)
	dummyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := dummyHead
	for i := 0; i < n; i++ {
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

// 链表排序适合归并排序
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid, prev := findMiddle(head)
	// 断开，分成两个链
	// 这儿有大坑，必须要保存mid前一个节点，从prev开始断。如果从mid开始断，当递归到链表长度为2时，会陷入死循环
	prev.Next = nil
	left := mergeSort(head)
	right := mergeSort(mid)
	return mergeList(left, right)
}

func mergeList(list1, list2 *ListNode) *ListNode {
	p := list1
	q := list2
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
			cur = cur.Next
		} else {
			cur.Next = q
			q = q.Next
			cur = cur.Next
		}
	}
	for p != nil {
		cur.Next = p
		p = p.Next
		cur = cur.Next
	}
	for q != nil {
		cur.Next = q
		q = q.Next
		cur = cur.Next
	}
	return dummyHead.Next
}

// 返回mid和prev
func findMiddle(head *ListNode) (*ListNode, *ListNode) {
	slow := head
	fast := head
	prev := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow, prev
}
