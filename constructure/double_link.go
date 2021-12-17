package constructure

type DLinkItem struct {
	val       int
	pre, next *DLinkItem
}

func (root *DLinkItem) Remove(item *DLinkItem) {
	item.pre.next = item.next
	if item.next != nil {
		item.next.pre = item.pre
	}
}

func (root *DLinkItem) Add(item *DLinkItem) {
	var p = root

	for p.next != nil {
		if p.next.val > item.val {
			p.next.pre = item
			item.next = p.next
			p.next = item
			item.pre = p
			return
		}
		p = p.next
	}

	p.next = item
	item.pre = p
}
