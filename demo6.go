package main

import "fmt"

type listNewNode struct {
	Data int
	Next *listNewNode
}

func main()  {

	list := []int{1,2,3,4}
	head := &listNewNode{Data: list[0]}
	tail := head

	for i := 1; i<len(list); i++  {
		tail.Next = &listNewNode{Data: list[i]}
		tail = tail.Next
	}
	//node4.Next = node5
	head.show()
	//mergerList := mergerTwoListNode(node1,node2)

	//fmt.Printf("This list has cycle? Yes or No: %v\n", mergerList)
}


func mergerTwoListNode(l1 *listNewNode, l2 *listNewNode) *listNewNode  {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data < l2.Data{
		l1.Next = mergerTwoListNode(l1.Next,l2)
		return l1
	}
	l2.Next = mergerTwoListNode(l2.Next,l1)
	return l2
}

func (l *listNewNode) show() {
	fmt.Println(l.Data)
	for l.Next != nil {
		l = l.Next
		fmt.Println(l.Data)
	}
}