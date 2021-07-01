package main

import "fmt"

type listNewNode struct {
	Data int
	Next *listNewNode
}

func main()  {

	a := new(listNewNode)
	a.Data = 1
	a.Next = &listNewNode{2, &listNewNode{4, nil}}

	b := new(listNewNode)
	b.Data = 1
	b.Next = &listNewNode{3, &listNewNode{4, nil}}
	c := mergerTwoListNode(a,b)
	c.show()
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