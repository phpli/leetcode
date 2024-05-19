package main

func main() {
	var head = new(newListNode)

}

type newListNode struct {
	data int
	next *newListNode
}

func revereNewList(head *newListNode) *newListNode {
	cur := head
	var pre *newListNode = nil
	for cur != nil {
		pre, cur, cur.next = cur.next, cur.next, pre
	}
	return pre
}
