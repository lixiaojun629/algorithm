package cache

//https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	cache    map[int]*DLinkNode
	link     *DoubleLinkList
	size     int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*DLinkNode),
		link:     NewDoubleLinkList(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.link.moveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		this.cache[key].Value = value
		this.link.moveToHead(node)
		return
	}

	node = &DLinkNode{
		Key:   key,
		Value: value,
	}
	this.cache[key] = node
	this.link.insertHead(node)
	this.size++
	if this.size > this.capacity {
		node := this.link.removeTail()
		delete(this.cache, node.Key)
		this.size--
	}
}

type DLinkNode struct {
	Key   int
	Value int
	Prev  *DLinkNode
	Next  *DLinkNode
}

type DoubleLinkList struct {
	headDummy *DLinkNode
	tailDummy *DLinkNode
}

func NewDoubleLinkList() *DoubleLinkList {
	head := &DLinkNode{Value: -1}
	tail := &DLinkNode{Value: -1}
	head.Next = tail
	tail.Prev = head
	return &DoubleLinkList{
		headDummy: head,
		tailDummy: tail,
	}
}

func (this *DoubleLinkList) remove(node *DLinkNode) *DLinkNode {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return node
}

func (this *DoubleLinkList) removeTail() *DLinkNode {
	return this.remove(this.tailDummy.Prev)
}

func (this *DoubleLinkList) insertHead(node *DLinkNode) {
	node.Next = this.headDummy.Next
	this.headDummy.Next.Prev = node

	this.headDummy.Next = node
	node.Prev = this.headDummy
}

func (this *DoubleLinkList) moveToHead(node *DLinkNode) {
	this.remove(node)
	this.insertHead(node)
}
