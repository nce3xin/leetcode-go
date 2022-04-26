// 239
package main

import (
	"container/heap"
	"fmt"
)

//Input: nums = [2,-2,-8,5,3,7,10], k = 3

//#Input: nums = [1,1,1,1,1,1,1], k = 1
//#Input: nums = [1,2,3,4,5,6,7], k = 5
func main() {
	nums := []int{2, -2, -8, 5, 3, 7, 10}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Printf("res: %v\n", res)
}

type Item struct {
	val   int
	index int
}

type IntHeap []Item

// 方法：就是构建一个最大堆
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0)
	maxHeap := IntHeap{}
	for i := 0; i < k; i++ {
		num := nums[i]
		heap.Push(&maxHeap, Item{val: num, index: i})
	}
	res = append(res, maxHeap[0].val)
	// 注意这里的条件，是i<n，因为要遍历剩下的所有数，我之前微软三面就卡在这个地方，脑子抽了，其实很简单的
	for i := k; i < n; i++ {
		heap.Push(&maxHeap, Item{val: nums[i], index: i})
		//这个地方直接通过 maxHeap[0].index 来判断将要弹出的元素的下标。不要通过heap.Pop()的返回值去判断
		if maxHeap[0].index <= i-k {
			heap.Pop(&maxHeap)
		}
		// 这里添加也是直接添加 maxHeap[0].val，不要通过heap.Pop()的返回值去添加
		res = append(res, maxHeap[0].val)
	}
	return res
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	if h[i].val > h[j].val {
		return true
	}
	return false
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *IntHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}
