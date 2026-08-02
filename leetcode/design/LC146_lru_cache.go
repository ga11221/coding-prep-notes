package main

type LRUCache struct {
	cache    map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    map[int]*Node{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		if node.Prev == nil { // the node is head
			return node.Val
		}
		prev := node.Prev
		next := node.Next
		if next == nil {
			this.tail = prev
		}
		prev.Next = next
		if next != nil {
			next.Prev = prev
		}
		node.Prev = nil
		this.head.Prev = node
		node.Next = this.head
		this.head = node
		return node.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		if node.Prev == nil { // the node is head
			return
		}
		prev := node.Prev
		next := node.Next
		if next == nil {
			this.tail = prev
		}
		prev.Next = next
		if next != nil {
			next.Prev = prev
		}
		node.Prev = nil
		this.head.Prev = node
		node.Next = this.head
		this.head = node
	} else {
		if len(this.cache) == this.capacity {
			temp := this.tail
			newTail := temp.Prev
			if newTail == nil {
				this.head = nil
				this.tail = nil
			} else {
				newTail.Next = nil
				this.tail = newTail
			}
			delete(this.cache, temp.Key)
		}
		node = &Node{
			Key: key,
			Val: value,
		}
		this.cache[key] = node
		if this.head == nil {
			this.head = node
			this.tail = node
			return
		}
		this.head.Prev = node
		node.Next = this.head
		this.head = node
	}
}
