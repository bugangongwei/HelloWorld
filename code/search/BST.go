package search

// 二叉排序树的实现

// BSTNode 二叉排序树的节点
type BSTNode struct {
	Index uint64
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(index uint64, left, right *BSTNode) *BSTNode {
	return &BSTNode{
		Index: index,
		Left:  left,
		Right: right,
	}
}

// BST 二叉排序树
type BST struct {
	Root *BSTNode
}

func NewBST(root *BSTNode) *BST {
	return &BST{Root: root}
}

// Search 查找下标为 index 的节点
func (bst *BST) Search(index uint64) (*BSTNode, error) {
	var fn func(root *BSTNode) (*BSTNode, error)

	fn = func(root *BSTNode) (*BSTNode, error) {
		if root == nil {
			return nil, ErrNotFound
		}

		if index < root.Index {
			return fn(root.Left)
		}

		if index > root.Index {
			return fn(root.Right)
		}

		return root, nil
	}

	return fn(bst.Root)
}

// Add 添加一个节点
func (bst *BST) Add(node *BSTNode) {
	var fn func(root *BSTNode)

	fn = func(root *BSTNode) {
		if root == nil || root.Index == node.Index {
			root = node
			return
		}

		if node.Index < root.Index {
			fn(root.Left)
		}

		if node.Index > root.Index {
			fn(root.Right)
		}
	}

	fn(bst.Root)
}

// Remove 删除一个节点
func (bst *BST) Remove(index uint64) error {
	var fn func(root *BSTNode) error

	fn = func(root *BSTNode) error {
		if root == nil {
			return ErrNotFound
		}

		if index < root.Index {
			fn(root.Left)
		}

		if index > root.Index {
			fn(root.Right)
		}

		if root.Left == nil {
			root = root.Right
		}

		if root.Right == nil {
			root = root.Left
		}

		root.Left.Right = root.Right
		return nil
	}

	return fn(bst.Root)
}
