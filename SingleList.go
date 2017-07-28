package main

import (
	"fmt"
)

//节点结构定义
type Node struct {
	data int
	next *Node
}

//在链表第一个位置插入元素
func HeadInsert(h, d *Node) {
	d.next = h.next
	h.next = d
}

//获取链表元素个数
func GetLength(h *Node) (n int) {
	for t := h.next; t != nil; t = t.next {
		n++
	}
	return
}

//获取链表第一个元素
func GetFirst(h *Node) *Node {
	return h.next
}

//获取链表最后一个元素
func GetLast(h *Node) *Node {
	var t *Node = h
	for t.next != nil {
		t = t.next
	}
	return t
}

//获取链表第n个元素
func GetElement(h *Node, n int) *Node {
	var t *Node = h
	for ; t.next != nil; t, n = t.next, n-1 {
		if n == 0 {
			break
		}
	}
	return t
}

/* 在指定位置插入新元素
 * h 链表头
 * e 新元素
 * n 插入位置
 * 返回布尔值表示是否成功插入
 */
func InsertElement(h, e *Node, n int) bool {
	if n <= 0 || n > GetLength(h) {
		return false
	}
	var i int = 0
	for t := h.next; t != nil; t = t.next {
		i++
		if n == i {
			e.next = t.next
			t.next = e
			break
		}
	}
	return true
}

/* 删除链表中某个元素
 * h 链表头
 * n 删除位置
 */
func DeleteElement(h *Node, n int) (*Node, bool) {
	if n <= 0 || n > GetLength(h) {
		return nil, false
	}
	var i int = 0
	var p *Node
	for t := h; t.next != nil; t = t.next {
		i++
		if i == n {
			p = t.next
			t.next = t.next.next
			p.next = nil
			break
		}
	}
	return p, true
}

/* 清空整个链表 */
func ClearList(h *Node) {
	var p, q *Node = h, nil
	for p.next != nil {
		q = p.next
		p = nil
		p = q
	}
	h.next = nil
}

func main() {
	var head, x Node
	x.data = 1000

	//头插法
	for j := 1; j <= 10; j++ {
		var b Node
		b.data = j * 10
		HeadInsert(&head, &b)
	}

	//在链表中间插入新元素
	InsertElement(&head, &x, GetLength(&head)/2)

	//尾部追加
	tmp := GetLast(&head)
	for i := 1; i <= 10; i++ {
		var d Node
		d.data = i
		tmp.next = &d
		tmp = &d
	}

	//删除中间的元素
	if v, ok := DeleteElement(&head, GetLength(&head)/2); ok {
		fmt.Println(v.data)
	}

	//清空链表
	ClearList(&head)
}

//一些关于链表的题目

//1.从尾到头打印单链表 (递归实现)
//      e 为链表第一个元素
func tailtohead(e *Node) {
	if e == nil {
		return
	}
	tailtohead(e.next)
	fmt.Printf("%d ", e.data)
}

//待补充
//2.删除一个无头单链表的非尾节点
//3.在无头单链表的一个节点前插入一个节点
//4.单链表实现约瑟夫环
//5.逆置 / 反转单链表
//6.单链表排序（冒泡排序&快速排序）
//7.合并两个有序链表, 合并后依然有序
//8.查找单链表的中间节点，要求只能遍历一次链表
//9.查找单链表的倒数第k个节点，要求只能遍历一次链表
