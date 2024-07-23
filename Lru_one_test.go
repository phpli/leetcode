package main

import "testing"

type OneNode struct {
	key, value int
	prev, next *OneNode
}

type OneDoubleLinkedList struct {
	head, tail *OneNode
}

func NewOneDoubleLinkedList() *OneDoubleLinkedList {
	return &OneDoubleLinkedList{}
}

func (ll *OneDoubleLinkedList) addOneNodeToFront(node *OneNode) {
	//设置当前节点next，肯定是ll 头节点
	//因为是加到头部node.pre 节点是nil
	node.next = ll.head
	node.prev = nil
	if ll.head != nil {
		ll.head.prev = node
	}
	ll.head = node
	if ll.tail == nil {
		ll.tail = node
	}
}

func (ll *OneDoubleLinkedList) removeOneEndNode() *OneNode {
	if ll.tail == nil {
		return nil
	}
	oneNode := ll.tail
	//判断原来尾节点的prev节点是否为空
	if ll.tail.prev != nil {
		//把prev的下一个节点next节点设置为空
		ll.tail.prev.next = nil
	} else {
		//只能证明 此链表原来就一个元素
		ll.head = nil
	}
	//赋值新的tail
	ll.tail = ll.tail.prev
	return oneNode
}

func (ll *OneDoubleLinkedList) moveOneNodeToFront(node *OneNode) {
	//如果node的节点就是链表的头节点，直接返回
	if node == ll.head {
		return
	}
	//断开该节点的节点连续
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == ll.tail {
		ll.tail = node.prev
	}
	//节点移动到头部
	node.next = ll.head
	node.prev = nil
	if ll.head != nil {
		ll.head.prev = node
	}

	ll.head = node
	if ll.tail == nil { //// 如果尾节点为空（链表之前是空的），将尾节点也指向新节点
		ll.tail = node
	}
}

type oneLruCache struct {
	capacity int
	dll      *OneDoubleLinkedList
	cache    map[int]*OneNode
}

func newOneLruCache(capacity int) *oneLruCache {
	return &oneLruCache{
		capacity: capacity,
		dll:      NewOneDoubleLinkedList(),
		cache:    make(map[int]*OneNode),
	}
}

func (c *oneLruCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.dll.moveOneNodeToFront(node)
		return node.value
	}
	return -1
}

func (c *oneLruCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.dll.moveOneNodeToFront(node)
		return
	}
	if len(c.cache) >= c.capacity {
		remove := c.dll.removeOneEndNode()
		delete(c.cache, remove.key)
	}
	oneNode := &OneNode{key: key, value: value}
	c.dll.addOneNodeToFront(oneNode)
	c.cache[key] = oneNode
}
func TestOneLRUCache(t *testing.T) {
	cache := newOneLruCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	if value := cache.Get(1); value != 1 {
		t.Errorf("Expected -1, got %d", value)
	}
	cache.Put(3, 3)
	if value := cache.Get(2); value != -1 {
		t.Errorf("Expected -1, got %d", value)
	}
	cache.Put(4, 4)
	if value := cache.Get(1); value != -1 {
		t.Errorf("Expected -1, got %d", value)
	}
	if value := cache.Get(3); value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
	if value := cache.Get(4); value != 4 {
		t.Errorf("Expected 4, got %d", value)
	}
}
