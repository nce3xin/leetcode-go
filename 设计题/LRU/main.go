package main

// https://leetcode-cn.com/problems/lru-cache/

// Node 双向链表节点
type Node struct {
	key, val int
	prev     *Node
	next     *Node
}

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	head, tail *Node
	size       int
}

type LRUCache struct {
	mapping  map[int]*Node
	cache    DoubleLinkedList
	capacity int
}

func (this *LRUCache) moveToHead(node *Node) {
	this.cache.removeNode(node)
	this.cache.insertToHead(node)
}

// node一定存在
func (this *DoubleLinkedList) removeNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	this.size--
}

func (this *LRUCache) insertToHead(node *Node) {
	this.cache.insertToHead(node)
}

func (this *DoubleLinkedList) insertToHead(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
	this.size++
}

func (this *LRUCache) setValue(node *Node, val int) {
	node.val = val
	return
}

func (this *LRUCache) getSize() int {
	return this.cache.size
}

func (this *LRUCache) removeFromTail() {
	node := this.cache.removeFromTail()
	// 需要将尾部元素对应的key从map中也删除
	delete(this.mapping, node.key)
}

func (this *DoubleLinkedList) removeFromTail() *Node {
	if this.head.next == this.tail {
		return nil
	}
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	LRU := LRUCache{
		capacity: capacity,
		mapping:  make(map[int]*Node),
		cache: DoubleLinkedList{
			head: &Node{0, 0, nil, nil},
			tail: &Node{0, 0, nil, nil},
			size: 0,
		},
	}
	LRU.cache.head.next = LRU.cache.tail
	LRU.cache.tail.prev = LRU.cache.head
	return LRU
}

func (this *LRUCache) Get(key int) int {
	// 如果不存在，返回-1
	// 如果存在，移动到头部，再返回key对应的val
	if _, ok := this.mapping[key]; !ok {
		return -1
	}
	node := this.mapping[key]
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	// 如果存在，移动到头部，再修改key对应的val
	// 如果不存在，先构造一个新Node，map中增加key。然后又分2种情况。1. cache已满，先删除尾部，再插入到头部。2. cache未满，直接插入到头部
	if node, ok := this.mapping[key]; ok {
		this.moveToHead(node)
		this.setValue(node, value)
		return
	}
	size := this.getSize()
	node := Node{
		key:  key,
		val:  value,
		prev: nil,
		next: nil,
	}
	if size == this.capacity {
		this.removeFromTail()
		// 别忘了在map中添加key的映射
		this.mapping[key] = &node
		this.insertToHead(&node)
	} else {
		// 别忘了在map中添加key的映射
		this.mapping[key] = &node
		this.insertToHead(&node)
	}
}
