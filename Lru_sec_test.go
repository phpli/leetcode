package main

import "testing"

type SecNode struct {
	key, value int
	prev, next *SecNode
}

type SecDoubleLinkedList struct {
	head, tail *SecNode
}

func NewSecDoubleLinkedList() *SecDoubleLinkedList {
	return &SecDoubleLinkedList{}
}

func (sl *SecDoubleLinkedList) AddSecNodeToFront(node *SecNode) {
	//先把当前数据设置好
	node.prev = nil
	node.next = sl.head
	if sl.head != nil {
		sl.head.prev = node //将当前头节点的 prev 指针指向新节点
	}
	sl.head = node      //// 将链表的 head 指针更新为新节点
	if sl.tail == nil { // 如果链表的 tail 为空，说明链表之前是空的
		sl.tail = node // 将链表的 tail 指针也更新为新节点
	}

}

func (sl *SecDoubleLinkedList) RemoveSecNodeEnd() *SecNode {
	if sl.tail == nil {
		return nil
	}
	node := sl.tail
	//断开连接
	if sl.tail.prev != nil {
		//判断尾节点的前节点不为空，现在要置为空，因为要删除
		sl.tail.prev.next = nil
	} else {
		//为空说明 只有一个节点
		sl.head = nil
	}
	//尾节点变成了 之前尾节点的prev节点
	sl.tail = sl.tail.prev
	return node
}

func (sl *SecDoubleLinkedList) MoveSecNodeToFront(node *SecNode) {
	if node == sl.head {
		return
	}
	//断开连接
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	//设置双向链表的尾节点
	if node == sl.tail {
		sl.tail = node.prev
	}
	//设置新的头节点
	node.prev = nil
	node.next = sl.head
	if sl.head != nil {
		sl.head.prev = node
	}

	//设置双向链表的头节点
	sl.head = node

	if sl.tail == nil {
		sl.tail = node
	}
}

type SecLruCache struct {
	capacity int
	hashMap  map[int]*SecNode
	list     *SecDoubleLinkedList
}

func NewSecLruCache(capacity int) *SecLruCache {
	return &SecLruCache{
		capacity: capacity,
		hashMap:  make(map[int]*SecNode),
		list:     NewSecDoubleLinkedList(),
	}
}

func (sl *SecLruCache) get(key int) int {
	if node, ok := sl.hashMap[key]; ok {
		sl.list.MoveSecNodeToFront(node)
		return node.value
	}
	return -1
}

func (sl *SecLruCache) put(key int, value int) {
	if node, ok := sl.hashMap[key]; ok {
		node.value = value
		sl.list.MoveSecNodeToFront(node)
		return
	}
	if len(sl.hashMap) >= sl.capacity {
		node := sl.list.RemoveSecNodeEnd()
		delete(sl.hashMap, node.key)
	}
	oneNode := &SecNode{key: key, value: value}
	sl.list.AddSecNodeToFront(oneNode)
	sl.hashMap[key] = oneNode
}

func TestSecLRUCache(t *testing.T) {
	cache := NewSecLruCache(2)
	cache.put(1, 1)
	cache.put(2, 2)

	if value := cache.get(1); value != 1 {
		t.Errorf("Expected -1, got %d", value)
	}
	cache.put(3, 3)
	if value := cache.get(2); value != -1 {
		t.Errorf("Expected -1, got %d", value)
	}
	cache.put(4, 4)
	if value := cache.get(1); value != -1 {
		t.Errorf("Expected -1, got %d", value)
	}
	if value := cache.get(3); value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
	if value := cache.get(4); value != 4 {
		t.Errorf("Expected 4, got %d", value)
	}
}
