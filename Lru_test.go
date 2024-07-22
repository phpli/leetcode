package main

import (
	"testing"
)

// 节点的结构体
type Node struct {
	key, value int
	prev, next *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

// 初始化双向链表
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// 添加节点到链表头部
func (list *DoublyLinkedList) AddToFront(node *Node) {
	//当前节点赋值
	node.next = list.head
	node.prev = nil
	if list.head != nil { //原来链表不是空值
		list.head.prev = node //原来
	}
	list.head = node
	if list.tail == nil {
		list.tail = node
	}
}

// 返回移除的节点
func (list *DoublyLinkedList) RemoveFromEnd() *Node {
	// 尾节点为空，直接返回
	if list.tail == nil {
		return nil
	}
	node := list.tail
	if list.tail.prev != nil {
		list.tail.prev.next = nil
	} else {
		list.head = nil
	}
	list.tail = list.tail.prev
	return node
}

// 移动某个节点到头部
func (list *DoublyLinkedList) MoveNodeToFrond(node *Node) {
	//判断是否为头节点
	if node == list.head {
		return
	}
	// 从链表中移除节点
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == list.tail {
		list.tail = node.prev
	}
	// 将节点移动到头部
	node.next = list.head
	node.prev = nil
	if list.head != nil {
		list.head.prev = node
	}
	list.head = node
	// 检查是否需要更新尾部指针
	if list.tail == nil {
		list.tail = node
	}
}

type LRUCache struct {
	capacity int
	hashmap  map[int]*Node
	list     *DoublyLinkedList
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		hashmap:  make(map[int]*Node),
		list:     NewDoublyLinkedList(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.hashmap[key]; ok {
		lru.list.MoveNodeToFrond(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, found := lru.hashmap[key]; found {
		node.value = value
		lru.list.MoveNodeToFrond(node)
		return
	}
	//fmt.Println(lru.list)
	if len(lru.hashmap) >= lru.capacity {
		nodeToRemove := lru.list.RemoveFromEnd()
		//fmt.Println(nodeToRemove)
		delete(lru.hashmap, nodeToRemove.key)
	}

	newNode := &Node{key: key, value: value}
	lru.list.AddToFront(newNode)
	lru.hashmap[key] = newNode
}

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	//fmt.Println(cache)
	cache.Put(1, 1)
	//fmt.Println(cache)
	cache.Put(2, 2)
	//fmt.Println(cache)
	if value := cache.Get(1); value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	cache.Put(3, 3) // 该操作会使得关键字 2 作废
	//fmt.Println(cache)
	if value := cache.Get(2); value != -1 {
		t.Errorf("Expected -1, got %d", value)
	}

	cache.Put(4, 4) // 该操作会使得关键字 1 作废
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
