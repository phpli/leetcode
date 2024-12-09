package main

import (
	"container/list"
	"fmt"
)

type LinkedHashMap struct {
	data     map[interface{}]*list.Element
	order    *list.List
	capacity int // 0 表示无限制
}

type pair struct {
	key   interface{}
	value interface{}
}

// NewLinkedHashMap 创建一个新的 LinkedHashMap
func NewLinkedHashMap(capacity int) *LinkedHashMap {
	return &LinkedHashMap{
		data:     make(map[interface{}]*list.Element),
		order:    list.New(),
		capacity: capacity,
	}
}

// Put 插入或更新键值对
func (lhm *LinkedHashMap) Put(key, value interface{}) {
	if elem, exists := lhm.data[key]; exists {
		// 如果键已存在，更新值并将其移动到末尾
		elem.Value.(*pair).value = value
		lhm.order.MoveToBack(elem)
	} else {
		// 如果键不存在，检查容量限制
		if lhm.capacity > 0 && lhm.order.Len() >= lhm.capacity {
			// 移除最老的键值对
			oldest := lhm.order.Front()
			delete(lhm.data, oldest.Value.(*pair).key)
			lhm.order.Remove(oldest)
		}
		// 插入新的键值对
		elem := lhm.order.PushBack(&pair{key, value})
		lhm.data[key] = elem
	}
}

// Get 获取值，并将键值对移动到末尾（访问顺序）
func (lhm *LinkedHashMap) Get(key interface{}) (interface{}, bool) {
	if elem, exists := lhm.data[key]; exists {
		// 将其移动到末尾
		lhm.order.MoveToBack(elem)
		return elem.Value.(*pair).value, true
	}
	return nil, false
}

// Remove 删除键值对
func (lhm *LinkedHashMap) Remove(key interface{}) bool {
	if elem, exists := lhm.data[key]; exists {
		delete(lhm.data, key)
		lhm.order.Remove(elem)
		return true
	}
	return false
}

// Keys 返回所有键的有序列表
func (lhm *LinkedHashMap) Keys() []interface{} {
	keys := make([]interface{}, 0, lhm.order.Len())
	for elem := lhm.order.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(*pair).key)
	}
	return keys
}

// Values 返回所有值的有序列表
func (lhm *LinkedHashMap) Values() []interface{} {
	values := make([]interface{}, 0, lhm.order.Len())
	for elem := lhm.order.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(*pair).value)
	}
	return values
}

// Size 返回当前大小
func (lhm *LinkedHashMap) Size() int {
	return lhm.order.Len()
}

// 示例使用
func main() {
	lhm := NewLinkedHashMap(3)

	lhm.Put("a", 1)
	lhm.Put("b", 2)
	lhm.Put("c", 3)

	fmt.Println("Keys:", lhm.Keys())     // ["a", "b", "c"]
	fmt.Println("Values:", lhm.Values()) // [1, 2, 3]

	lhm.Get("a")    // 访问 "a"
	lhm.Put("d", 4) // 插入 "d"，"b" 被移除（容量限制）

	fmt.Println("Keys:", lhm.Keys())     // ["c", "a", "d"]
	fmt.Println("Values:", lhm.Values()) // [3, 1, 4]
}
