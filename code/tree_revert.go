package code

// 二叉树翻转, 简单迭代写法
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var (
		s    = new(StackImp)
		node *TreeNode
	)

	s.Push(root)

	for s.GetLen() != 0 {
		node = s.Pop()

		var tmp = node.Left
		node.Left = node.Right
		node.Right = tmp

		if node.Left != nil {
			s.Push(node.Left)
		}

		if node.Right != nil {
			s.Push(node.Right)
		}
	}
	return root
}

// 二叉树翻转, 统一前中后序的迭代写法(这根本没必要, 但是可以复习一下这个点的掌握程度)
// 中序[左中右]
func InvertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var (
		s    = new(StackImp)
		node *TreeNode
	)

	s.Push(root)

	for s.GetLen() != 0 {
		node = s.Pop()

		if node != nil {
			if node.Right != nil {
				s.Push(node.Right)
			}

			s.Push(node)
			s.Push(nil)

			if node.Left != nil {
				s.Push(node.Left)
			}
		} else {
			node = s.Pop()

			var tmp = node.Left
			node.Left = node.Right
			node.Right = tmp
		}

	}

	return root
}

// 翻转二叉树, 层序遍历写法
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	q := new(Queue)
	q.Push(root)

	for q.Size() != 0 {
		node := q.Pop().(*TreeNode)

		var tmp = node.Left
		node.Left = node.Right
		node.Right = tmp

		if node.Left != nil {
			q.Push(node.Left)
		}

		if node.Right != nil {
			q.Push(node.Right)
		}
	}

	return root
}
