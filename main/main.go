package main

import (
	"bugangongwei/HelloWorld/verification"
)

func main() {
	// 704 标准二分问题
	// fmt.Println(code.Search([]int{1, 3, 5, 6}, 0))
	// 704 plus 二分加插入
	// fmt.Println(code.SearchOrInsert([]int{1, 3, 5, 6}, 0))
	// 27 数组删除元素
	// fmt.Println(code.RemoveElement([]int{3, 2, 2, 3}, 3))
	// 977 有序数组的平方、
	// fmt.Println(code.SortedSquares([]int{-4, 0, 1, 3, 10}))
	// 209 长度最小的子数组
	// fmt.Println(code.MinSubArrayLen(1, []int{1, 4, 4}))
	// 59 螺旋打印数组
	// fmt.Println(code.GenerateMatrix(3))
	// 203 移除地铁元素
	// head = [1,2,6,3,4,5,6], val = 6
	// head = [1,6], val  = 6
	// head = [1,2], val = 6
	// head = [1], val = 6 done
	// head = [], val=6 done
	// head = [6, 1, 2], val = 6 done
	// head = [6], val = 6 done
	// p6 := code.ListNode{Val: 6, Next: nil}
	// p5 := code.ListNode{Val: 5, Next: &p6}
	// p4 := code.ListNode{Val: 4, Next: &p5}
	// p3 := code.ListNode{Val: 1, Next: nil}
	// p2 := code.ListNode{Val: 2, Next: &p3}
	// p1 := code.ListNode{Val: 2, Next: &p2}
	// p0 := code.ListNode{Val: 1, Next: &p1}

	// nHead := code.RemoveElementsV2(&p0, 2)
	// fmt.Println(nHead)

	// p := nHead
	// for p != nil {
	// 	fmt.Println(p.Val)
	// 	p = p.Next
	// }

	// 707 设计链表
	// p4 := code.ListNode{Val: 5, Next: nil}
	// p3 := code.ListNode{Val: 4, Next: &p4}
	// p2 := code.ListNode{Val: 8, Next: &p3}
	// p1 := code.ListNode{Val: 1, Next: &p2}
	// p0 := code.ListNode{Val: 4, Next: &p1}

	// pf := code.ListNode{Val: 5, Next: nil}
	// pe := code.ListNode{Val: 4, Next: &pf}
	// pd := code.ListNode{Val: 8, Next: &pe}
	// pc := code.ListNode{Val: -4, Next: nil}
	// pb := code.ListNode{Val: 0, Next: &pc}
	// pa := code.ListNode{Val: 2, Next: nil}

	// 206 反转链表
	// newHead := code.ReverseList(&p0)

	// 24 两两交换链表中的节点
	// newHead := code.SwapPairs(&p0)

	// 19 删除倒数第 n 个链表元素
	// newHead := code.RemoveNthFromEnd(&p0, 2)

	// 20.07 链表相交
	// newHead := code.GetIntersectionNode(&p0, &pa)
	// p := newHead
	// for p != nil {
	// 	fmt.Println(p.Val)
	// 	p = p.Next
	// }

	// 142 环形链表找入环口
	// head := &code.ListNode{Val: 1, Next: &pa}
	// pa.Next = head
	// enter := code.DetectCycle(nil)
	// fmt.Println(verification.Return_err(1))

	// fmt.Println("the end")

	/* 二叉树 */
	// root := &code.TreeNode{Val: 1}
	// root.Left = &code.TreeNode{Val: 2}
	// root.Right = &code.TreeNode{Val: 3}
	// root.Left.Left = &code.TreeNode{Val: 4}

	// 前序, 中序, 后序遍历
	// arr := code.PostorderTraversal2(root)
	// fmt.Println(arr)

	// root := &code.Node{Val: 3}
	// root.Left = &code.Node{Val: 9}
	// root.Right = &code.Node{Val: 20}
	// root.Left.Left = &code.Node{Val: 4}
	// root.Left.Right = &code.Node{Val: 5}
	// root.Right.Left = &code.Node{Val: 15}
	// root.Right.Right = &code.Node{Val: 7}
	// code.Connect2(root)

	// root := &code.TreeNode{Val: 4}
	// root.Left = &code.TreeNode{Val: 2}
	// root.Right = &code.TreeNode{Val: 7}
	// root.Left.Left = &code.TreeNode{Val: 1}
	// root.Left.Right = &code.TreeNode{Val: 3}
	// root.Right.Left = &code.TreeNode{Val: 6}
	// root.Right.Right = &code.TreeNode{Val: 9}

	// 层序遍历
	// root := &code.TreeNode{Val: 3}
	// root.Left = &code.TreeNode{Val: 9}
	// root.Right = &code.TreeNode{Val: 20}
	// root.Right.Left = &code.TreeNode{Val: 15}
	// root.Right.Right = &code.TreeNode{Val: 7}
	// fmt.Println(code.MinDepth(root))

	// 回溯
	verification.UnmarshalToRaw()
}
