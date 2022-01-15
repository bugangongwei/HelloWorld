package model

import "fmt"

type LRUCache struct {
	capacity int
	root     *DRLinkItem // 双向环形链表的 root
	cache    map[int]*DRLinkItem
}

func Constructor(capacity int) LRUCache {
	root := &DRLinkItem{}
	root.next = root
	root.pre = root

	return LRUCache{capacity: capacity, root: root, cache: make(map[int]*DRLinkItem, capacity)}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if ok {
		this.root.moveToFront(v)
		return v.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	v, ok := this.cache[key]

	// 插入
	if !ok {
		if len(this.cache) >= this.capacity {
			// 删除最后一个
			old := this.root.removeOldest()
			fmt.Println("删除最后一个元素", old)
			delete(this.cache, old.key)
		}

		newV := &DRLinkItem{val: value, key: key}
		fmt.Println("添加元素", newV)
		this.root.insertNewest(newV)
		this.cache[key] = newV
		return
	}

	// 更新
	this.cache[key].val = value
	this.root.moveToFront(v)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
