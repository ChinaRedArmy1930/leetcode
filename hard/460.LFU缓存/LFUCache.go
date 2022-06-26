package main

import (
	"container/list"
	"fmt"
)

var (
	debug = false
)

type ListNodeLFU struct {
	key       int
	value     int
	frequency int
	prev      *ListNodeLFU
	next      *ListNodeLFU
}

type LFUCache struct {
	minFre   int //最小频率
	count    int //LFU长度
	capacity int //LFU容量
	nodeMap  map[int]*ListNodeLFU
	freMap   map[int]*list.List
	nodeHead *ListNodeLFU
	nodeTail *ListNodeLFU
}

//将新节点插入到LFU的开头
func (this *LFUCache) addNew2Head(new *ListNodeLFU) {
	this.count++
	this.nodeMap[new.key] = new

	if this.nodeHead == nil {
		this.nodeHead = new
		this.nodeTail = new
		return
	}

	new.next = this.nodeHead
	new.prev = nil
	this.nodeHead.prev = new
	this.nodeHead = new
}

func (this *LFUCache) debug() {
	if !debug {
		return
	}
	fmt.Print("{")
	for l := this.nodeHead; l != nil; l = l.next {
		fmt.Printf("[%d, %d] c(%d) ", l.key, l.value, l.frequency)
	}
	fmt.Printf("minfre := %d ", this.minFre)
	fmt.Println("}")
}

func (this *LFUCache) delNode(v *ListNodeLFU) {
	this.count--
	delete(this.nodeMap, v.key)
	if v == this.nodeHead {
		this.nodeHead = this.nodeHead.next
		if this.nodeHead != nil {
			this.nodeHead.prev.next = nil
			this.nodeHead.prev = nil
		}
		return
	}

	if v == this.nodeTail {
		this.nodeTail = this.nodeTail.prev
		if this.nodeTail != nil {
			this.nodeTail.next.prev = nil
			this.nodeTail.next = nil
		}
		return
	}

	v.prev.next = v.next
	v.next.prev = v.prev
	v.prev = nil
	v.next = nil

}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		minFre:   0,
		count:    0,
		capacity: capacity,
		nodeMap:  map[int]*ListNodeLFU{},
		freMap:   map[int]*list.List{},
		nodeHead: nil,
		nodeTail: nil,
	}
}

// Get 如果存在LFU中, 直接返回
func (this *LFUCache) Get(key int) int {
	if v, ok := this.nodeMap[key]; ok {
		//修改freList
		for l := this.freMap[v.frequency].Front(); l != nil; l = l.Next() {
			if l.Value.(*ListNodeLFU) == v {
				this.freMap[v.frequency].Remove(l)
			}

		}

		if this.freMap[v.frequency].Len() == 0 && v.frequency == this.minFre {
			this.minFre = v.frequency + 1
		}

		v.frequency++

		if this.freMap[v.frequency] == nil {
			this.freMap[v.frequency] = list.New()
		}

		this.freMap[v.frequency].PushFront(v)
		this.debug()
		return v.value
	}

	this.debug()
	return -1
}

// Put 如果没有超过容量,且是链表中没有的数据 直接插入到链表头部
//     如果没有超过容量,且是链表中有的数据 修改frequency和value
//如果超过容量, 则删除用的最少的那个,  如果有多个相同的, 删除时间最远的那个
func (this *LFUCache) Put(key int, value int) {
	if v, ok := this.nodeMap[key]; !ok {
		if this.count >= this.capacity {
			if this.freMap[this.minFre] != nil {
				//找到frequency最小的节点, 然后删除
				this.delNode(this.freMap[this.minFre].Back().Value.(*ListNodeLFU))
				this.freMap[this.minFre].Remove(this.freMap[this.minFre].Back())
			} else {
				this.debug()
				return
			}
		}

		new := &ListNodeLFU{
			key:       key,
			value:     value,
			frequency: 1,
			prev:      nil,
			next:      nil,
		}
		this.addNew2Head(new)

		if this.freMap[1] == nil {
			this.freMap[1] = list.New()
		}

		this.freMap[1].PushFront(new)

		this.minFre = 1
		this.debug()
		return
	} else {
		//修改频率表
		freList := this.freMap[v.frequency]
		for l := freList.Front(); l != nil; l = l.Next() {
			if l.Value.(*ListNodeLFU) == v {
				this.freMap[v.frequency].Remove(l)
				break
			}
		}
		if this.freMap[v.frequency].Len() == 0 && v.frequency == this.minFre {
			this.minFre = v.frequency + 1
		}

		v.frequency++

		v.value = value
		if this.freMap[v.frequency] == nil {
			this.freMap[v.frequency] = list.New()
		}
		this.freMap[v.frequency].PushFront(v)
	}
	this.debug()
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	/*
				obj := Constructor(2)
				obj.Put(1, 1)
				obj.Put(2, 2)           //[{1,1} c(1),  {2,2} c(1)]
				fmt.Println(obj.Get(1)) //[{1,1} c(2),  {2,2} c(1)]
				obj.Put(3, 3)           //[{1,1} c(2),  {3,3} c(1)]
				fmt.Println(obj.Get(2)) //-1
				fmt.Println(obj.Get(3))
				obj.Put(4, 4)
				fmt.Println(obj.Get(1))
				fmt.Println(obj.Get(3))
				fmt.Println(obj.Get(4))


			obj := Constructor(0)
			obj.Put(0, 0)
			obj.Get(0)


		obj3 := Constructor(1)
		obj3.Put(2, 1)
		fmt.Println(obj3.Get(2))
		obj3.Put(3, 2)
		fmt.Println(obj3.Get(2))
		fmt.Println(obj3.Get(3))
	*/
	/*
		["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
		[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	*/

	/*
		[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]

	*/
	obj4 := Constructor(10)
	obj4.Put(10, 13)
	obj4.Put(3, 17)
	obj4.Put(6, 11)
	obj4.Put(10, 5)           //{[10,5] c(2), [3,17] c(1), 6,11 c(1)}
	obj4.Put(9, 10)           //{[10,5] c(2), [3,17] c(1), 6,11 c(1), [9,10] c(1)}
	fmt.Println(obj4.Get(13)) // -1
	obj4.Put(2, 19)           //{[10,5] c(2), [3,17] c(1), 6,11 c(1), [9,10] c(1), [2,19] c(1)}
	fmt.Println(obj4.Get(2))  //{[10,5] c(2), [3,17] c(1), 6,11 c(1), [9,10] c(1), [2,19] c(2)}
	fmt.Println(obj4.Get(3))  //{[10,5] c(2), [3,17] c(2), 6,11 c(1), [9,10] c(1), [2,19] c(2)}
	obj4.Put(5, 25)           //{[10,5] c(2), [3,17] c(2), 6,11 c(1), [9,10] c(1), [2,19] c(2),[5,25] c(1)}
	//OK
	fmt.Println(obj4.Get(8))
	obj4.Put(9, 22)
	obj4.Put(5, 5)
	obj4.Put(1, 30)
	fmt.Println(obj4.Get(11))
	obj4.Put(9, 12)
	fmt.Println(obj4.Get(7))
	fmt.Println(obj4.Get(5))
	fmt.Println(obj4.Get(8))
	fmt.Print("obj4.Get(9) => ")
	fmt.Println(obj4.Get(9)) // 12
	obj4.Put(4, 30)
	obj4.Put(9, 3)
	fmt.Print("obj4.Get(9) => ")
	fmt.Println(obj4.Get(9))
	fmt.Print("obj4.Get(10) => ")
	fmt.Println(obj4.Get(10))
	fmt.Print("obj4.Get(10) => ")
	fmt.Println(obj4.Get(10))
	obj4.Put(6, 14)
	obj4.Put(3, 1)
	fmt.Print("obj4.Get(3) => ")
	fmt.Println(obj4.Get(3))
	obj4.Put(10, 11)
	fmt.Print("obj4.Get(8) => ")
	fmt.Println(obj4.Get(8))
	obj4.Put(2, 14)
	fmt.Print("obj4.Get(1) => ")
	fmt.Println(obj4.Get(1))
	fmt.Print("obj4.Get(5) => ")
	fmt.Println(obj4.Get(5))
	fmt.Print("obj4.Get(4) => ")
	fmt.Println(obj4.Get(4))
	obj4.Put(11, 4)
	obj4.Put(12, 24)
	obj4.Put(5, 18)
	fmt.Print("obj4.Get(13) => ")
	fmt.Println(obj4.Get(13))
	obj4.Put(7, 23)
	fmt.Print("obj4.Get(8) => ")
	fmt.Println(obj4.Get(8))
	fmt.Print("obj4.Get(12) => ")
	fmt.Println(obj4.Get(12))
	obj4.Put(3, 27)
	obj4.Put(2, 12)
	fmt.Print("obj4.Get(5) => ")
	fmt.Println(obj4.Get(5))
	obj4.Put(2, 9)
	obj4.Put(13, 4)
	obj4.Put(8, 18)
	obj4.Put(1, 7)
	fmt.Print("obj4.Get(6) => ")
	fmt.Println(obj4.Get(6))
	obj4.Put(9, 29)
	obj4.Put(8, 21)
	fmt.Print("obj4.Get(5) => ")
	fmt.Println(obj4.Get(5))
	obj4.Put(6, 30)
	obj4.Put(1, 12)
	fmt.Print("obj4.Get(10) => ")
	fmt.Println(obj4.Get(10))
	obj4.Put(4, 15)
	obj4.Put(7, 22)
	obj4.Put(11, 26)
	obj4.Put(8, 17)
	obj4.Put(9, 29)
	fmt.Print("obj4.Get(5) => ")
	fmt.Println(obj4.Get(5))
	obj4.Put(3, 4)
	obj4.Put(11, 30)
	fmt.Println(obj4.Get(12))
	obj4.Put(4, 29)
	fmt.Print("obj4.Get(3) => ")
	fmt.Println(obj4.Get(3))
	fmt.Print("obj4.Get(9) => ")
	fmt.Println(obj4.Get(9))
	fmt.Print("obj4.Get(6) => ")
	fmt.Println(obj4.Get(6))
	obj4.Put(3, 4)
	fmt.Print("obj4.Get(1) => ")
	fmt.Println(obj4.Get(1))
	fmt.Print("obj4.Get(10) => ")
	fmt.Println(obj4.Get(10))
	obj4.Put(3, 29)
	obj4.Put(10, 28)
	obj4.Put(1, 20)
	obj4.Put(11, 13)
	fmt.Print("obj4.Get(3) => ")
	fmt.Println(obj4.Get(3))
	obj4.Put(3, 12)
	obj4.Put(3, 8)
	obj4.Put(10, 9)
	obj4.Put(3, 26)
	fmt.Println(obj4.Get(8))
	fmt.Println(obj4.Get(7)) //出错
	fmt.Println(obj4.Get(5))
	obj4.Put(13, 17)
	obj4.Put(2, 27)
	obj4.Put(11, 15)
	fmt.Println(obj4.Get(12))
	obj4.Put(9, 19)
	obj4.Put(2, 15)
	obj4.Put(3, 16)
	fmt.Println(obj4.Get(1))
	obj4.Put(12, 17)
	obj4.Put(9, 1)
	obj4.Put(6, 19)
	fmt.Println(obj4.Get(4))
	fmt.Println(obj4.Get(5))
	fmt.Println(obj4.Get(5))
	obj4.Put(8, 1)
	obj4.Put(11, 7)
	obj4.Put(5, 2)
	obj4.Put(9, 28)
	fmt.Println(obj4.Get(1))
	obj4.Put(2, 2)
	obj4.Put(7, 4)
	obj4.Put(4, 22)
	obj4.Put(7, 24)
	obj4.Put(9, 26)
	obj4.Put(13, 28)
	obj4.Put(11, 26)
}
