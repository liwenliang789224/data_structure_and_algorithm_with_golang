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

/* 插入元素
 * e 为被插入元素
 * at 为插入位置
 */
func (l *Node) InsertElement(e, at *Node) {
	e.next = at.next
	at.next.prev = e
	at.next = e
	e.prev = at
}

//删除元素
func (l *Node) DeleteElement(e *Node) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
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

func main() {
	strings := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelveth", "thirteenth", "fourteenth", "fifteenth", "sixteenth", "seventeenth", "eighteenth", "nineteenth", "twentieth"}
	list := New(len(strings))
	t := list
	for _, v := range strings {
		t.data = v
		t = t.next
	}
	fmt.Println(list.Len())
	for i, x := 0, list; i < list.Len(); i, x = i+1, x.next {
		fmt.Printf("%s ", x.data)
	}
}
