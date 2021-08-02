package code

/*
环形链表II:
slow: 以步幅1的速度前进的指针
fast: 以步幅2的速度前速的指针
x: 从开始到入环点的距离
y: 从入环点到相遇点的距离
z: 相遇点回到入环点的距离
在环中相遇时, slow指针已经走了 x+y+m(y+z), fast指针已经走了 x+y+n(y+z)
(x+y+m(y+z))(2/1) = x+y+n(y+z)
2(x+y)+2m(y+z) = (x+y)+n(y+z)
x+y=(n-2m)(y+z)
x = (n-2m-1)(y+z)+z
所以: x和z之间, 仅仅隔着整数个环
指针a从走x, 指针b走环, 则两者一定会相遇
*/
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// 寻找相遇点
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}
