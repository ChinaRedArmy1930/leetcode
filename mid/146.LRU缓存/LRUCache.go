package main

import (
	"fmt"
	"sync"
)

type ListNodeLRU struct {
	Key   int
	Value int
	Prev  *ListNodeLRU
	Next  *ListNodeLRU
}

type LRUCache struct {
	mu       sync.Mutex
	Count    int //当前容量
	Capacity int //最大容量
	NodeMap  map[int]*ListNodeLRU
	NodeHead *ListNodeLRU
	NodeTail *ListNodeLRU
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		NodeMap:  make(map[int]*ListNodeLRU, 0),
		NodeHead: nil,
		NodeTail: nil,
	}
}

// Get 如果不存在, 则返回-1, 如果存在,则将其提到第一位
func (this *LRUCache) Get(key int) int {
	if v, ok := this.NodeMap[key]; !ok {
		return -1
	} else {
		nodeAddress := v
		//本来就是第一个
		if nodeAddress.Prev == nil {
			return nodeAddress.Value
		}

		if nodeAddress == this.NodeTail {
			this.NodeTail = nodeAddress.Prev
		}

		nodeAddress.Prev.Next = nodeAddress.Next
		if nodeAddress.Next != nil {
			nodeAddress.Next.Prev = nodeAddress.Prev
		}

		nodeAddress.Prev = nil
		nodeAddress.Next = this.NodeHead
		this.NodeHead.Prev = nodeAddress
		this.NodeHead = nodeAddress

		return nodeAddress.Value
	}
}

// Put 插入到当前队列的头部
// 如果超过了容量, 则将最后一个删除, 并且将当前的插入到队列的头部
// 如果key存在, 则修改其内容,并且提到第一个
func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if v, ok := this.NodeMap[key]; ok {
		nodeAddress := v
		nodeAddress.Value = value
		//本来就是第一个
		if nodeAddress.Prev == nil {
			return
		}

		if nodeAddress == this.NodeTail {
			this.NodeTail = nodeAddress.Prev
		}

		nodeAddress.Prev.Next = nodeAddress.Next
		if nodeAddress.Next != nil {
			nodeAddress.Next.Prev = nodeAddress.Prev
		}

		nodeAddress.Prev = nil
		nodeAddress.Next = this.NodeHead
		this.NodeHead.Prev = nodeAddress
		this.NodeHead = nodeAddress
		return
	}

	newNode := &ListNodeLRU{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	this.NodeMap[key] = newNode

	//第一个节点
	if this.NodeHead == nil {
		this.NodeHead = newNode
		this.NodeTail = newNode
		this.Count++
		return
	}

	newNode.Next = this.NodeHead
	this.NodeHead.Prev = newNode
	this.NodeHead = newNode
	this.Count++

	//节点数小于容量
	if this.Count <= this.Capacity {
		return
	}

	//节点数大于容量
	delete(this.NodeMap, this.NodeTail.Key)
	this.NodeTail = this.NodeTail.Prev
	this.Count--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	/*
		obj := Constructor(2)
		obj.Put(1, 1)           //{[1,1]}
		obj.Put(2, 2)           //{[2,2],[1,1]}
		fmt.Println(obj.Get(1)) //{[1,1],[2,2]} => 1
		obj.Put(3, 3)           //{[3,3],[1,1]}
		fmt.Println(obj.Get(2)) //   => -1
		obj.Put(4, 4)           //{[4,4][3,3]}
		fmt.Println(obj.Get(1)) //=> -1
		fmt.Println(obj.Get(3)) // {[3,3],[4,4]} => 3
		fmt.Println(obj.Get(4))
	*/

	obj2 := Constructor(2)
	obj2.Put(2, 1)           //{[2,1]}
	obj2.Put(1, 1)           //{[1,1],[2,1]}
	obj2.Put(2, 3)           //{[2,3],[1,1]}
	obj2.Put(4, 1)           //{[4,1],[2,3]}
	fmt.Println(obj2.Get(1)) // -1
	fmt.Println(obj2.Get(2)) // 3
}
