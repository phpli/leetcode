package main

type listMidNode struct {
	Data int
	Next *listMidNode
}

func main() {

	head := &listMidNode{}
	listOne := head
	listOne.Data = 1
	listOne.Next = &listMidNode{}
	listOne = listOne.Next
	listOne.Data = 1

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
