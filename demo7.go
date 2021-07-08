package main

import "fmt"

/**
delete list 倒数第N个节点
*/

type listDeleteNode struct {
	Data int
	Next *listDeleteNode
}

func main() {
	head := &listDeleteNode{}
	head.Data = 1
	head.Next = &listDeleteNode{}
	head.Next.Data = 2
	head.Next.Next = &listDeleteNode{}
	head.Next.Next.Data = 3
	head.Next.Next.Next = &listDeleteNode{}
	head.Next.Next.Next.Data = 4
	//ret := head
	other := head
	//for ret != nil {
	//	fmt.Print(ret.Data, " ")
	//	ret = ret.Next
	//}
	//
	//ret = first(head,1)
	//for ret != nil {
	//	fmt.Print(ret.Data, " ")
	//	ret = ret.Next
	//}
	//
	other = second(other, 1)
	for other != nil {
		fmt.Print(other.Data, " ")
		other = other.Next
	}
}

/**
循环
*/
func first(head *listDeleteNode, n int) *listDeleteNode {
	var h, p, pre *listDeleteNode
	h = new(listDeleteNode) ///在头节点之前放置一个节点
	h.Next = head
	var l int
	pre = h
	p = head
	/*第一次遍历，获得链表的长度 L*/
	for {
		if p != nil {
			l++
			p = p.Next
		} else {
			break
		}
	}
	var i int
	p = head
	/*第二次遍历，到达顺数的 L-n+1 个节点处*/
	for i = 1; i < l-n+1; i++ {
		p = p.Next
		pre = pre.Next
	}
	pre.Next = p.Next
	return h.Next
}

/**
fast slow
*/
func second(head *listDeleteNode, n int) *listDeleteNode {
	fast, slow := head, head
	//slow = head
	//slow_pre = h
	//快指针先走
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}

	// 将快指针和慢指针同时向后移动，直到快指针移到了链表的最后一个结点
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//删除 节点
	slow.Next = slow.Next.Next
	return head
}
