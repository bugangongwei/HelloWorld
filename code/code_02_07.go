package code

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB

	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}

	return curA
}
