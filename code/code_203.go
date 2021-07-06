package code

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	var p = head
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

// 使用虚拟头节点
func RemoveElementsV2(head *ListNode, val int) *ListNode {
	vitualHead := &ListNode{Next: head}

	p := vitualHead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return vitualHead.Next
}
