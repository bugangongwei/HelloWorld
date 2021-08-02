package code

// 二叉树的最小深度
func MinDepth(root *TreeNode) int {
	return minDepthReverse(root)
}

func minDepthReverse(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right == nil {
		return 1 + minDepthReverse(root.Left)
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepthReverse(root.Right)
	}

	return 1 + minHigth(minDepthReverse(root.Left), minDepthReverse(root.Right))
}

func minHigth(x, y int) int {
	if x < y {
		return x
	}
	return y
}
