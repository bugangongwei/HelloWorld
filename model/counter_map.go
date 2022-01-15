package model

import "fmt"

type AllOne struct {
	m    map[string]*DRLinkItemWithMap
	root *DRLinkItemWithMap
}

/** Initialize your data structure here. */
func AllOneConstructor() AllOne {
	root := &DRLinkItemWithMap{}
	root.next = root
	root.pre = root
	root.keys = make(map[string]struct{})

	return AllOne{
		m:    make(map[string]*DRLinkItemWithMap),
		root: root,
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	item, ok := this.m[key]

	if ok {
		this.m[key] = this.root.AddOne(item, key)
		fmt.Println("after update ", this.m[key], this.root.pre, this.root.next)
	} else {
		this.m[key] = this.root.AddOne(this.root, key)
		fmt.Println("after insert ", this.m[key], this.root.pre, this.root.next)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	item, ok := this.m[key]

	if ok {
		// item.val--
		// fmt.Println("入参....", item)
		this.root.RemoveItem(item, key)
	}
	fmt.Println("after delete ", item, item.next, this.root, this.root.pre, this.root.next)
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	for k, _ := range this.root.pre.keys {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	for k, _ := range this.root.next.keys {
		return k
	}
	return ""
}
