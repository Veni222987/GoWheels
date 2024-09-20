package main

type doubleListNode struct {
	key   int
	val   int
	left  *doubleListNode
	right *doubleListNode
}

type LRUCache struct {
	cap  int
	sMap map[int]*doubleListNode
	head *doubleListNode
	tail *doubleListNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		sMap: make(map[int]*doubleListNode),
		cap:  capacity,
		head: &doubleListNode{},
		tail: &doubleListNode{},
	}
	l.head.right, l.tail.left = l.tail, l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.sMap[key]; ok {
		v.left.right, v.right.left = v.right, v.left
		this.head.right, this.head.right.left, v.left, v.right = v, v, this.head, this.head.right
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.sMap[key]; ok {
		v.val = value //刷新值
		// 刷新位置
		v.left.right, v.right.left = v.right, v.left
		this.head.right, this.head.right.left, v.left, v.right = v, v, this.head, this.head.right
		return
	}
	this.sMap[key] = &doubleListNode{
		key:   key,
		val:   value,
		left:  this.head,
		right: this.head.right,
	}
	this.head.right.left = this.sMap[key]
	this.head.right = this.sMap[key]

	if len(this.sMap) > this.cap {
		rmk := this.tail.left.key
		// 删除尾部元素
		this.tail.left.left.right = this.tail
		this.tail.left = this.tail.left.left
		delete(this.sMap, rmk)
	}
}
