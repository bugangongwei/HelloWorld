package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSkipList 单测
func TestSkipList(t *testing.T) {
	// head: 拥有最高的层数, 没有值
	head := NewSkipListNode(0, "head", DefaultMaxLevel)
	// currentLevel 从 1 开始
	sl := NewSkipList(head, 1)

	// 添加节点
	sl.Add(NewSkipListNode(1, 1, 0))
	sl.Add(NewSkipListNode(2, 2, 0))

	// 查询节点
	node, err := sl.Search(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, node.Val.(int))

	// 修改节点
	err = sl.Update(2, 200)
	assert.NoError(t, err)

	// 查询修改后的节点
	node, err = sl.Search(2)
	assert.NoError(t, err)
	assert.Equal(t, 200, node.Val.(int))

	// 删除节点
	sl.Remove(2)

	// 查询删除后的节点
	node, err = sl.Search(2)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, node)
}