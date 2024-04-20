package main

import "fmt"

/**
\ 反转单链表
那么这道题其实就是把指针指向前一个节点
*/

type ListNode struct {
	data int
	Next *ListNode
}

// 置换位置
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		//pre, cur, cur.Next = cur, cur.Next, pre //
		nextTemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTemp
	}
	return pre
}

func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.data = i
		cur = cur.Next
	}
}

func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, "")
	}
	fmt.Println()
}

func main() {
	head := &ListNode{data: 1, Next: &ListNode{data: 2, Next: &ListNode{data: 3, Next: &ListNode{data: 4, Next: nil}}}}
	//CreateNode(head, 10)
	PrintNode("前：", head)
	yyy := reverseList(head)
	PrintNode("后：", yyy)

}
