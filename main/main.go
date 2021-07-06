package main

import (
	"bugangongwei/HelloWorl/code"
	"fmt"
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
	p2 := code.MyLinkedList{Val: 3, Next: nil}
	p1 := code.MyLinkedList{Val: 2, Next: &p2}
	p0 := code.MyLinkedList{Val: 1, Next: &p1}

	head := p0
	head.AddAtHead(0)
	fmt.Println(head.Val)

	p := &head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

}
