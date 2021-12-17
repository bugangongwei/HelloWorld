package constructure

type DRLinkItem struct {
	pre, next *DRLinkItem
	val       int
	key       int
}

func (root *DRLinkItem) moveToFront(item *DRLinkItem) {
	// 从链表中删除 item
	item.next.pre = item.pre
	item.pre.next = item.next
	// 把 item 插入到 root 后面
	item.next = root.next
	root.next.pre = item
	item.pre = root
	root.next = item
}

func (root *DRLinkItem) insertNewest(item *DRLinkItem) {
	// item 插入到 root 后面
	item.next = root.next
	root.next.pre = item
	item.pre = root
	root.next = item
}

func (root *DRLinkItem) removeOldest() *DRLinkItem {
	tmp := root.pre
	root.pre.pre.next = root
	root.pre = root.pre.pre

	return tmp
}

type DRLinkItemWithMap struct {
	pre, next *DRLinkItemWithMap
	val       int
	keys      map[string]struct{}
}

// item 存在并加1
func (root *DRLinkItemWithMap) AddOne(item *DRLinkItemWithMap, key string) *DRLinkItemWithMap {
	// delete from old counter
	defer func() {
		// fmt.Println("判断 ", item != root, item, root)
		if item != root {
			root.RemoveItem(item, key)
		}
	}()

	// merge to next
	if item.next.val == item.val+1 {
		item.next.keys[key] = struct{}{}
		return item.next
	}

	// insert between item and item.next

	newItem := &DRLinkItemWithMap{val: item.val + 1, keys: map[string]struct{}{key: {}}}
	// fmt.Println(">>>", item.next)
	newItem.next = item.next
	item.next.pre = newItem
	item.next = newItem
	newItem.pre = item

	// fmt.Println("===> ", newItem, newItem.pre, newItem.next, item.pre, item.next)

	return newItem
}

func (root *DRLinkItemWithMap) RemoveItem(item *DRLinkItemWithMap, key string) {
	if _, ok := item.keys[key]; ok {
		if len(item.keys) == 1 && item.val == 1 {
			// remove
			item.pre.next = item.next
			item.next.pre = item.pre
		} else if len(item.keys) == 1 && item.val > 1 {
			item.val--
		} else {
			// remove key from item
			delete(item.keys, key)
			// new val
			newVal := item.val - 1
			if newVal == 0 {
				return
			}
			// new item with new val
			newItem := &DRLinkItemWithMap{val: newVal, keys: map[string]struct{}{key: {}}}
			// new item merge or new item insert
			if newItem.val == item.pre.val {
				// merge to pre
				item.pre.keys[key] = struct{}{}
			} else {
				// insert between and item
				newItem.pre = item.pre
				newItem.next = item.next
				item.pre.next = newItem
				item.next.pre = newItem
			}
		}
	}
}
