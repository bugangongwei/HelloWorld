package code

// [1, 2, 3, null]
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	var pre *ListNode
	pre = nil

	for p != nil {
		current := p
		p = p.Next

		current.Next = pre
		pre = current
	}
	return pre
}
