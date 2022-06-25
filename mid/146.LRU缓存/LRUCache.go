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

//将一个新的节点添加到链表的头部
func (this *LRUCache) addNew2Head(new *ListNodeLRU) {
	if new == nil {
		return
	}

	//注意map别忘记加了
	if _, ok := this.NodeMap[new.Key]; ok {
		return
	}

	this.NodeMap[new.Key] = new
	this.Count++

	//链表为空
	if this.NodeHead == nil {
		this.NodeHead = new
		this.NodeTail = new
		return
	}

	//链表不为空
	new.Next = this.NodeHead
	new.Prev = nil
	this.NodeHead.Prev = new
	this.NodeHead = new
}

//将某个存在key提升到链表的头部
func (this *LRUCache) makeOldRecently(node *ListNodeLRU) {
	if node == nil {
		return
	}

	//注意, 如果节点是链表尾部 则需要处理是否是唯一一个节点的情况
	if node == this.NodeTail {
		if node.Prev == nil {
			this.NodeTail = this.NodeHead
		} else {
			this.NodeTail = node.Prev
		}
	}

	//如果这个链表就在链表头部, 则什么都不做
	if node == this.NodeHead {
		return
	}

	//node的前后两个节点链到一起
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	//将node初始化为一个新的节点, 然后添加到链表头
	this.Count--
	delete(this.NodeMap, node.Key)
	node.Prev = nil
	node.Next = nil

	this.addNew2Head(node)
}

//删除末尾的节点
func (this *LRUCache) delTailNode() {
	if this.NodeTail == nil {
		return
	}

	this.Count--
	delete(this.NodeMap, this.NodeTail.Key)

	if this.NodeTail.Prev != nil {
		this.NodeTail = this.NodeTail.Prev
		this.NodeTail.Next = nil
	} else {
		this.NodeTail = nil
		this.NodeHead = nil
	}
}

// Get 如果不存在, 则返回-1, 如果存在,则将其提到第一位
func (this *LRUCache) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()
	if v, ok := this.NodeMap[key]; ok {
		this.makeOldRecently(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) dump() {
	tmp := this.NodeHead
	for tmp != nil {
		fmt.Print("[")
		fmt.Print(tmp.Key)
		fmt.Print(", ")
		fmt.Print(tmp.Value)
		fmt.Print("] ")
		tmp = tmp.Next
	}
	fmt.Println()
}

// Put 插入到当前队列的头部
// 如果超过了容量, 则将最后一个删除, 并且将当前的插入到队列的头部
// 如果key存在, 则修改其内容,并且提到第一个
func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if v, ok := this.NodeMap[key]; ok {
		v.Value = value
		this.makeOldRecently(v)
		return
	}

	this.addNew2Head(&ListNodeLRU{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	})

	if this.Count > this.Capacity {
		this.delTailNode()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

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

	obj2 := Constructor(2)
	obj2.Put(2, 1)           //{[2,1]}
	obj2.Put(1, 1)           //{[1,1],[2,1]}
	obj2.Put(2, 3)           //{[2,3],[1,1]}
	obj2.Put(4, 1)           //{[4,1],[2,3]}
	fmt.Println(obj2.Get(1)) // -1
	fmt.Println(obj2.Get(2)) // 3

	obj3 := Constructor(1)
	obj3.Put(2, 1)
	fmt.Println(obj3.Get(2))

	obj4 := Constructor(1)
	obj4.Put(2, 1)
	fmt.Println(obj4.Get(2))
	obj4.Put(3, 2)
	fmt.Println(obj4.Get(2))
	fmt.Println(obj4.Get(3))

	obj5 := Constructor(3)
	obj5.Put(1, 1)
	obj5.Put(2, 2)
	obj5.Put(3, 3)
	obj5.Put(4, 4) //{[4,4],[3,3],[2,2]}

	fmt.Println(obj5.Get(4)) // 4
	fmt.Println(obj5.Get(3)) // 3 {[3,3],[4,4],[2,2]}
	fmt.Println(obj5.Get(2)) // 2 {[2,2],[3,3],[4,4]}
	fmt.Println(obj5.Get(1)) // -1

	obj5.Put(5, 5)           //{[5,5],[2,2],[3,3]}
	fmt.Println(obj5.Get(1)) // -1
	fmt.Println(obj5.Get(2)) // {[2,2],[5,5],[3,3]} => 2
	fmt.Println(obj5.Get(3)) //{[3,3],[2,2],[5,5] ==> 3
	fmt.Println(obj5.Get(4)) // -1
	fmt.Println(obj5.Get(5)) //{[5,5],[3,3],[2,2]} ==> 5
	 
	obj6 := Constructor(10)
	obj6.Put(10, 13)          //{[10,13]}
	obj6.Put(3, 17)           //{[3,17], [10,13]}
	obj6.Put(6, 11)           //{[6,11], [3,17], [10,13]}
	obj6.Put(10, 5)           //{[10,5],[6,11], [3,17]}
	obj6.Put(9, 10)           //{[9,10],[10,5],[6,11], [3,17]}
	fmt.Println(obj6.Get(13)) // -1
	obj6.Put(2, 19)           //{[2,19], [9,10],[10,5],[6,11], [3,17]}
	fmt.Println(obj6.Get(2))  // 19
	fmt.Println(obj6.Get(3))  //17 {[3,17], [2,19], [9,10],[10,5],[6,11]}
	obj6.Put(5, 25)           //{[5,25],[3,17], [2,19], [9,10],[10,5],[6,11]}
	fmt.Println(obj6.Get(8))  // -1
	obj6.Put(9, 22)           //{[9,22],[5,25],[3,17], [2,19], [10,5],[6,11]}
	obj6.Put(5, 5)            //{[5,5],[9,22],[3,17], [2,19], [10,5],[6,11]}
	obj6.Put(1, 30)           //{[1,30],[5,5],[9,22],[3,17], [2,19], [10,5],[6,11]}
	fmt.Println(obj6.Get(11)) // -1
	obj6.Put(9, 12)           //{[9,12],[1,30],[5,5],[3,17], [2,19], [10,5],[6,11]}
	fmt.Println(obj6.Get(7))  // -1
	fmt.Println(obj6.Get(5))  // 5 {[5,5],[9,12],[1,30],[3,17], [2,19], [10,5],[6,11]}
	fmt.Println(obj6.Get(8))  // -1
	fmt.Println(obj6.Get(9))  // 12  {[9,12],[5,5],[1,30],[3,17], [2,19], [10,5],[6,11]}
	obj6.Put(4, 30)           // {[4,30],[9,12],[5,5],[1,30],[3,17], [2,19], [10,5],[6,11]}
	obj6.Put(9, 3)            // {[9,3],[4,30],[5,5],[1,30],[3,17], [2,19], [10,5],[6,11]}
	fmt.Println(obj6.Get(9))  //3
	fmt.Println(obj6.Get(10)) //5 {[5,5],[9,3],[4,30],[1,30],[3,17], [2,19], [10,5],[6,11]}
	fmt.Println(obj6.Get(10)) //5 {[10,5],[5,5],[9,3],[4,30],[1,30],[3,17], [2,19], [6,11]}
	obj6.Put(6, 14)           // -1
	obj6.Put(3, 1)            //30 {[1,30],[10,5],[5,5],[9,3],[4,30],[3,17], [2,19], [6,11]}
	fmt.Println(obj6.Get(3))  //17 {[3,17],[1,30],[10,5],[5,5],[9,3],[4,30], [2,19], [6,11]}
	obj6.Put(10, 11)          //
	fmt.Println(obj6.Get(8))  //-1
}
