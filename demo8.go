package main

import "fmt"

type listMidNode struct {
	Data int
	Next *listMidNode
}

func main() {

	list := []int{1, 2, 3, 4, 5}
	head := &listMidNode{Data: list[0]}
	tail := head
	for i := 1; i < len(list); i++ {
		tail.Next = &listMidNode{Data: list[i]}
		tail = tail.Next
	}
	midUp := middleUpNode(head)
	midDown := middleDownNode(head)
	fmt.Printf("%d", midUp.Data)
	fmt.Printf("%d", midDown.Data)
}

//1 2 3 4 5 6 7 8
/**
1.输入链表头节点，奇数长度返回中点，偶数长度返回上中点 。
*/
func middleUpNode(head *listMidNode) *listMidNode {
	pre := &listMidNode{} //虚拟头节点
	pre.Data = 1          //补一个头数据
	pre.Next = head
	slow := pre
	fast := pre
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//2.输入链表头节点，奇数长度返回中点，偶数长度返回下中点 。这道题是leetcode上的第876道题，叫【链表的中间节点】。
//快慢指针
// 2 3 4 5 6 7 8
func middleDownNode(head *listMidNode) *listMidNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
