package main

// 23
import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (lnp ListNodeHeap) Len() int {
	return len(lnp)
}

func (lnp ListNodeHeap) Less(i, j int) bool {
	if lnp[i].Val < lnp[j].Val {
		return true
	}
	return false
}

func (lnp ListNodeHeap) Swap(i, j int) {
	lnp[i], lnp[j] = lnp[j], lnp[i]
}

func (lnp *ListNodeHeap) Push(x interface{}) {
	*lnp = append(*lnp, x.(*ListNode))
}

func (lnp *ListNodeHeap) Pop() interface{} {
	item := (*lnp)[len(*lnp)-1]
	*lnp = (*lnp)[:len(*lnp)-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := dummy
	minHeap := ListNodeHeap{}
	// 将 k 个链表的头结点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(&minHeap, head)
		}
	}
	for len(minHeap) > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
		p.Next = node
		p = p.Next
	}
	return dummy.Next
}
