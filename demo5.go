package main

import "fmt"

type ListCycleNode struct {
	Data interface{}
	Next *ListCycleNode
}

func HasCycle(head *ListCycleNode) bool  {
	if head == nil{
		return false
	}
	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return  false
}

func main() {
	var node1 = new(ListCycleNode)
	var node2 = new(ListCycleNode)
	var node3 = new(ListCycleNode)
	var node4 = new(ListCycleNode)
	var node5 = new(ListCycleNode)

	node1.Data = 1
	node2.Data = 2
	node3.Data = 3
	node4.Data = 4
	node5.Data = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	//node4.Next = node5

	hasCycle := HasCycle(node1)

	fmt.Printf("This list has cycle? Yes or No: %v\n", hasCycle)
}


