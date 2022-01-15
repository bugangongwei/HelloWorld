package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSkipList 单测
func TestSkipList(t *testing.T) {
	head := NewSkipListNode(0, "head", DefaultMaxLevel)
	sl := NewSkipList(head, 1)

	sl.Add(NewSkipListNode(1, 1, 0))
	sl.Add(NewSkipListNode(2, 2, 0))

	node, err := sl.Search(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, node.Val.(int))

	err = sl.Update(2, 200)
	assert.NoError(t, err)

	node, err = sl.Search(2)
	assert.NoError(t, err)
	assert.Equal(t, 200, node.Val.(int))

	sl.Remove(2)

	node, err = sl.Search(2)
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, node)
}