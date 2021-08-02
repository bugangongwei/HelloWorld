package code

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Arr []int
}

/*
二叉树前序, 中序, 后序遍历
递归解法
*/
func preorderTraversal(root *TreeNode) []int {
	res := &Result{
		Arr: make([]int, 0),
	}
	preTraversal(root, res)
	return res.Arr
}

func preTraversal(root *TreeNode, res *Result) {
	if root == nil {
		return
	}

	res.Arr = append(res.Arr, root.Val)
	preTraversal(root.Left, res)
	preTraversal(root.Right, res)
}

/*
二叉树前序, 中序, 后序遍历
迭代解法
*/
// 前序
// 中入栈, 中出栈 [中]
// 右入栈
// 左入栈
// 左出栈 [左]
// 右出栈 [右]
func PreTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := new(StackImp)
	stack.Push(root)

	for stack.GetLen() != 0 {
		node := stack.Pop()
		result = append(result, node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return result
}

// 中序
func InorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	stack := new(StackImp)
	cur := root

	// 为空, 出栈一个元素, 往右子树走
	// 不为空, 往左子树走
	// 栈为空的时候, 如果右子树不为空, 则右子树入栈
	// 左子树为空的时候, 如果栈不为空, 则需要出栈
	for stack.GetLen() != 0 || cur != nil {
		if cur == nil {
			fmt.Println(stack.GetLen(), cur)
			cur = stack.Pop()
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			stack.Push(cur)
			cur = cur.Left
		}
	}

	return result
}

// 后序 [左右中]
// 前序+换左右顺序 [中右左]
// result 翻转 [左右中]
func PostorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := new(StackImp)
	stack.Push(root)

	for stack.GetLen() != 0 {
		node := stack.Pop()
		result = append(result, node.Val)

		if node.Left != nil {
			stack.Push(node.Left)
		}

		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	reverse(result)

	return result
}

func reverse(result []int) {
	i, j := 0, len(result)-1

	for i < j {
		tmp := result[i]
		result[i] = result[j]
		result[j] = tmp
		i++
		j--
	}
}

type StackImp struct {
	arr []*TreeNode
	len int
}

func (s *StackImp) Top() *TreeNode {
	if s == nil {
		panic("pop: nil stack")
	}

	if len(s.arr) == 0 {
		panic("pop from empty stack")
	}

	return s.arr[(len(s.arr) - 1):][0]
}

func (s *StackImp) Pop() *TreeNode {
	if s == nil {
		panic("pop: nil stack")
	}

	if len(s.arr) == 0 {
		panic("pop from empty stack")
	}

	s.len--
	res := s.arr[(len(s.arr) - 1):][0]
	s.arr = s.arr[:(len(s.arr) - 1)]
	return res
}

func (s *StackImp) Push(top *TreeNode) {
	if s == nil {
		panic("push: nil stack")
	}

	s.arr = append(s.arr, top)
	s.len++
}

func (s *StackImp) GetLen() int {
	return s.len
}
