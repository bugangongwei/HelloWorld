package code

import (
	"fmt"
	"log"
)

/*
层序遍历相关
*/

// 层序遍历
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	q := new(Queue)
	q.Push(root)

	for q.Size() != 0 {
		size := q.Size()
		tmp := make([]int, 0)

		for i := 0; i < size; i++ {
			node := q.Pop().(*TreeNode)
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		result = append(result, tmp)
	}

	return result
}

type Queue struct {
	arr []interface{}
}

func (q *Queue) Push(node interface{}) {
	q.arr = append(q.arr, node)
}

func (q *Queue) Pop() interface{} {
	if len(q.arr) == 0 {
		log.Fatal("pop from empty queue")
	}

	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}

func (q *Queue) Top() interface{} {
	if len(q.arr) == 0 {
		log.Fatal("top from empty queue")
	}

	return q.arr[0]
}

func (q *Queue) Size() int {
	return len(q.arr)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
func Connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := new(Queue)

	q.Push(root)
	for q.Size() != 0 {
		size := q.Size()
		fmt.Println("size", size)
		preNode := q.Top().(*Node)

		for i := 0; i < size; i++ {
			node := q.Pop().(*Node)
			fmt.Println("出", node.Val, q.Size())
			if node.Left != nil {
				q.Push(node.Left)
				fmt.Println("入", node.Left.Val, q.Size())
			}
			if node.Right != nil {
				q.Push(node.Right)
				fmt.Println("入", node.Right.Val, q.Size())
			}

			if i > 0 {
				preNode.Next = node
				fmt.Println(preNode.Val, "->", node.Val)
				preNode = node
			}
		}
	}

	return root
}

// 117. 填充每个节点的下一个右侧节点指针 II
func Connect2(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}

// 662. 二叉树最大宽度
func WidthOfBinaryTree(root *TreeNode) int {
	var (
		q        = new(FullNodeQueue)
		result   = 0
		curDepth = 0
		left     = 0
	)
	q.Push(&FullNode{root, 0, 0})

	for q.Size() > 0 {
		s := q.Size()
		for i := 0; i < s; i++ {
			node := q.Pop()
			if node.node != nil {
				q.Push(&FullNode{node.node.Left, node.depth + 1, node.position * 2})
				q.Push(&FullNode{node.node.Right, node.depth + 1, node.position*2 + 1})
				if curDepth != node.depth {
					curDepth = node.depth
					left = node.position
				}
				result = max(result, node.position-left+1)
			}
		}
	}

	return result
}

type FullNode struct {
	node     *TreeNode
	depth    int
	position int
}

type FullNodeQueue struct {
	arr []*FullNode
}

func (q *FullNodeQueue) Pop() *FullNode {
	if len(q.arr) == 0 {
		log.Fatal("cannot pop from empty queue")
	}

	node := q.arr[0]
	q.arr = q.arr[1:]
	return node
}

func (q *FullNodeQueue) Top() *FullNode {
	if len(q.arr) == 0 {
		log.Fatal("empty queue")
	}

	return q.arr[0]
}

func (q *FullNodeQueue) Push(node *FullNode) {
	q.arr = append(q.arr, node)
}

func (q *FullNodeQueue) Size() int {
	return len(q.arr)
}
