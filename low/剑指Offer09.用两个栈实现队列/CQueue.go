package main

import "sync"

//todo
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type Stack struct {
	mu sync.Mutex
	N  *Node
}

func (s Stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := -1
	if s.N.Next != nil {
		result = s.N.Next.Val
	}
	s.N.Next.Next = nil
	s.N.Next.Prev = nil
	s.N.Next = s.N.Next.Next

	if s.N.Next.Next != nil {
		s.N.Next.Next.Prev = s.N

	}

	return result

}

type CQueue struct {
	S1 Stack
	S2 Stack
}

func Constructor() CQueue {
	return CQueue{
		S1: Stack{
			mu: sync.Mutex{},
			N: &Node{
				Val:  0,
				Next: nil,
				Prev: nil,
			},
		},
		S2: Stack{
			mu: sync.Mutex{},
			N: &Node{
				Val:  0,
				Next: nil,
				Prev: nil,
			},
		},
	}
}

func (this *CQueue) AppendTail(value int) {
	value = 1
}

func (this *CQueue) DeleteHead() int {
	return 1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
