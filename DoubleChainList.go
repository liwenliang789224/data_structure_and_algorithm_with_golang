package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

//链表元素个数
func (l *Node) Len() (n int) {
	if l != nil {
		n = 1
		for p := l.next; p != l; p = p.next {
			n++
		}
	}
	return
}

//创建一个新的拥有n个元素的循环链表
func New(n int) *Node {
	if n <= 0 {
		return nil
	}
	r := new(Node)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Node{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}