package code

// head = [1,2,3,4] [A, 2, 1, 3(h), 4]
// head = [] []
// head = [1] [1]
func SwapPairs(head *ListNode) *ListNode {
	var (
		h           = head                  // [两两]中的第一个元素
		virtualHead = &ListNode{Next: head} // 虚拟头节点
		pre         = virtualHead           // 上一组[两两]中的 tail
	)

	for h != nil {
		if h.Next == nil {
			break
		}

		// swap
		cur := h
		h = h.Next.Next // h 后移
		pre.Next = cur.Next
		pre.Next.Next = cur
		cur.Next = h
		pre = cur
	}

	return virtualHead.Next
}
