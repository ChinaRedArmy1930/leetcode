package main

import (
	"fmt"
)

const N = 100010

type MyLinkedList struct {
	head  int
	tail  int
	idx   int
	count int
	E     [N]int
	NE    [N]int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		idx:  0,
		head: -1,
		tail: -1,
		E:    [N]int{},
		NE:   [N]int{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	i := 0
	j := 0
	if index >= this.idx {
		return -1
	}

	for i = this.head; i != -1; i = this.NE[i] {
		if j == index {
			break
		}
		j++
	}

	if i == -1 {
		return -1
	}
	return this.E[i]
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.E[this.idx] = val
	this.NE[this.idx] = this.head
	this.head = this.idx
	if this.tail == -1 {
		this.tail = this.idx
	}
	this.idx++
	this.count++
}

func (this *MyLinkedList) AddAtTail(val int) {
	i := 0

	if this.idx == 0 {
		this.AddAtHead(val)
		return
	}

	for i = this.head; this.NE[i] != -1; i = this.NE[i] {
	}

	this.E[this.idx] = val
	this.NE[i] = this.idx
	this.NE[this.idx] = -1
	this.idx++
	this.count++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.count {
		this.AddAtTail(val)
		return
	}

	i := 0
	j := 0
	p := 0
	for i = this.head; i != -1; i = this.NE[i] {
		//fmt.Printf("%d => %d \n", i, this.E[i])
		if j == index {
			break
		}
		j++
		p = i
	}

	if i == -1 {
		return
	} else {
		this.E[this.idx] = val
		this.NE[this.idx] = i

		this.NE[p] = this.idx
		this.idx++
		this.NE[this.idx] = -1
	}
	this.count++
}

func (this *MyLinkedList) dump() {
	fmt.Print("dump: ")
	j := 0
	for i := this.head; i != -1; i = this.NE[i] {
		fmt.Printf("[%d]=>%d ", j, this.E[i])
		j++
	}
	fmt.Println()

	for i := 0; i < this.idx; i++ {
		//fmt.Printf("%d ", this.E[i])
	}
	//fmt.Println()
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.NE[this.head]
		this.count--
		return
	}

	i := 0
	j := 0
	p := 0
	for i = this.head; i != -1 && j != index; i = this.NE[i] {
		j++
		p = i
	}

	//没找到
	if i == -1 {
		return
	}

	//倒数第一个
	if this.NE[i] == -1 {
		this.NE[p] = -1
	} else {
		this.NE[p] = this.NE[this.NE[p]]
	}
	this.count--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	/*
		["MyLinkedList",obj.AddAtHead(),obj.AddAtIndex()obj.AddAtTail()obj.AddAtTail()obj.AddAtTail()obj.AddAtIndex()obj.AddAtTail()obj.AddAtHead(),obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.AddAtIndex()obj.AddAtTail()fmt.Println(obj.Get())fmt.Println(obj.Get())obj.AddAtHead(),obj.AddAtTail()obj.AddAtTail()fmt.Println(obj.Get())obj.AddAtTail()obj.AddAtTail()obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.AddAtHead(),obj.AddAtTail()obj.AddAtIndex()fmt.Println(obj.Get())obj.AddAtTail()obj.AddAtIndex()obj.AddAtHead(),obj.AddAtTail()obj.AddAtIndex()fmt.Println(obj.Get())obj.AddAtHead(),obj.AddAtTail()obj.AddAtIndex()obj.AddAtHead(),obj.AddAtIndex()obj.AddAtTail()obj.AddAtHead(),obj.AddAtIndex()obj.AddAtTail()obj.AddAtHead(),obj.DeleteAtIndex(),fmt.Println(obj.Get())obj.AddAtIndex()fmt.Println(obj.Get())obj.AddAtIndex()obj.AddAtTail()obj.AddAtTail()fmt.Println(obj.Get())obj.DeleteAtIndex(),fmt.Println(obj.Get())obj.AddAtHead(),obj.AddAtTail()obj.AddAtIndex()obj.AddAtIndex()obj.AddAtIndex()obj.AddAtHead(),obj.AddAtTail()obj.AddAtIndex()obj.DeleteAtIndex(),obj.AddAtHead(),obj.AddAtHead(),obj.AddAtTail()fmt.Println(obj.Get())obj.AddAtTail()obj.AddAtIndex()obj.AddAtHead(),obj.DeleteAtIndex(),obj.AddAtHead(),obj.DeleteAtIndex(),fmt.Println(obj.Get())fmt.Println(obj.Get())obj.AddAtTail()obj.AddAtIndex()fmt.Println(obj.Get())obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.AddAtHead(),obj.AddAtHead(),obj.AddAtIndex()fmt.Println(obj.Get())obj.AddAtTail()obj.AddAtHead(),obj.AddAtIndex()fmt.Println(obj.Get())obj.AddAtHead(),obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.DeleteAtIndex(),obj.AddAtHead(),obj.AddAtTail()fmt.Println(obj.Get())obj.AddAtHead(),obj.AddAtTail()obj.AddAtHead(),obj.AddAtHead(),obj.DeleteAtIndex(),fmt.Println(obj.Get())obj.AddAtHead()]
		[[],[55],[1,90],[51],[91],[12],[2,72],[17],[82],[4],[7],[7],[5,75],[54],[6],[2],[8],[35],[36],[10],[40],[43],[12],[3],[78],[89],[3,41],[10],[96],[5,37],[51],[26],[16,91],[18],[11],[66],[22,20],[44],[17,16],[95],[2],[14,2],[99],[51],[1],[11],[22,99],[20],[25,42],[72],[45],[2],[4],[32],[55],[84],[32,64],[26,14],[30,80],[88],[51],[27,71],[15],[8],[60],[37],[25],[96],[25,53],[36],[8],[85],[42],[20],[34],[78],[42,76],[26],[30],[39],[27],[93],[19,75],[8],[24],[32],[25,98],[21],[95],[18],[45],[24],[38],[8],[20],[83],[71],[78],[55],[29],[11],[84]]
	*/
	obj.AddAtHead(55)
	obj.AddAtIndex(1, 90)
	obj.AddAtTail(51)
	obj.AddAtTail(91)
	obj.AddAtTail(12)
	obj.AddAtIndex(2, 72)
	obj.AddAtTail(17)
	obj.AddAtHead(82)
	obj.DeleteAtIndex(4)
	obj.DeleteAtIndex(7)
	obj.DeleteAtIndex(7)
	obj.AddAtIndex(5, 75)
	obj.AddAtTail(54)
	fmt.Println(obj.Get(6))
	fmt.Println(obj.Get(2))
	obj.AddAtHead(8)
	obj.AddAtTail(35)
	obj.AddAtTail(36)
	fmt.Println(obj.Get(10))
	obj.AddAtTail(40)
	obj.AddAtTail(43)
	obj.DeleteAtIndex(12)
	obj.DeleteAtIndex(3)
	obj.AddAtHead(78)
	obj.AddAtTail(89)
	obj.AddAtIndex(3, 41)
	fmt.Println(obj.Get(10))
	obj.AddAtTail(96)
	obj.AddAtIndex(5, 37)
	obj.AddAtHead(51)
	obj.AddAtTail(26)
	obj.AddAtIndex(16, 91)
	fmt.Println(obj.Get(18))
	obj.AddAtHead(11)
	obj.AddAtTail(66)

	obj.dump()

	obj.AddAtIndex(22, 20)
	obj.dump()

	obj.AddAtHead(44)
	obj.AddAtIndex(17, 16)
	obj.AddAtTail(95)
	obj.AddAtHead(2)
	obj.AddAtIndex(14, 2)
	obj.AddAtTail(99)
	obj.AddAtHead(51)
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(11))
	obj.AddAtIndex(22, 99)
	fmt.Println(obj.Get(20))
	obj.AddAtIndex(25, 42)
	obj.AddAtTail(72)
	obj.AddAtTail(45)
	fmt.Println(obj.Get(2))
	obj.DeleteAtIndex(4)
	fmt.Println(obj.Get(32))
	obj.AddAtHead(55)
	obj.AddAtTail(84)
	obj.AddAtIndex(32, 64)
	obj.AddAtIndex(26, 14)
	obj.AddAtIndex(30, 80)
	obj.AddAtHead(88)
	obj.AddAtTail(51)
	obj.AddAtIndex(27, 71)
	obj.DeleteAtIndex(15)
	obj.AddAtHead(8)
	obj.AddAtHead(60)
	obj.AddAtTail(37)
	fmt.Println(obj.Get(25))
	obj.AddAtTail(96)
	obj.AddAtIndex(25, 53)
	obj.AddAtHead(36)
	obj.DeleteAtIndex(8)
	obj.AddAtHead(85)
	obj.DeleteAtIndex(42)
	fmt.Println(obj.Get(20))

	//
	fmt.Println("====")
	obj.dump()
	fmt.Println(obj.Get(34))
	obj.AddAtTail(78)
	obj.AddAtIndex(43, 76)
	fmt.Println(obj.Get(26))
	obj.DeleteAtIndex(30)
	obj.DeleteAtIndex(39)
	obj.AddAtHead(27)
	obj.AddAtHead(93)
	obj.AddAtIndex(19, 75)
	fmt.Println(obj.Get(8))
	obj.AddAtTail(24)
	obj.AddAtHead(32)
	obj.AddAtIndex(25, 98)
	fmt.Println(obj.Get(21))
	obj.AddAtHead(95)
	obj.DeleteAtIndex(18)
	obj.DeleteAtIndex(45)
	obj.DeleteAtIndex(24)
	obj.AddAtHead(38)
	obj.AddAtTail(8)
	fmt.Println(obj.Get(20))
	obj.AddAtHead(83)
	obj.AddAtTail(71)
	obj.AddAtHead(78)
	obj.AddAtHead(55)
	obj.DeleteAtIndex(29)
	fmt.Println(obj.Get(11))
	obj.AddAtHead(84)
}
