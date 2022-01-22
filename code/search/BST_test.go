package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	// 创建一棵新的二叉排序树
	bst := NewBST(nil)

	// 添加子树
	bst.Add(NewBSTNode(5, nil, nil))
	bst.Add(NewBSTNode(1, nil, nil))
	bst.Add(NewBSTNode(7, nil, nil))

	// 查询添加的节点
	node, err := bst.Search(7)
	assert.NoError(t, err)
	assert.Equal(t, 7, node.Index)

	// 删除节点
	err = bst.Remove(5)
	assert.NoError(t, err)

	// 查询被删除的节点
	node, err = bst.Search(5)
	assert.Error(t, ErrNotFound, err)
	assert.Nil(t, node)
}
