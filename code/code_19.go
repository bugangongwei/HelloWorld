package code

/*
输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5] 示例 2：
输入：head = [1], n = 1 输出：[] 示例 3：
输入：head = [1,2], n = 1 输出：[1]
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	vHead := &ListNode{0, head}
	target, pre := head, vHead

	// 拉开快慢指针的差距, n
	for i := 0; i < n; i++ {
		target = target.Next
	}

	for ; target != nil; target = target.Next {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return vHead.Next

}
