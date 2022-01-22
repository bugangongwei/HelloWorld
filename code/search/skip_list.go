package search

import (
	"errors"
	"math/rand"
)

/*
	reference: https://z.itpub.net/article/detail/14191D414D97A8CCC558B09FF5DB0BA1
	reference: https://weread.qq.com/web/reader/d35323e0597db0d35bd957bk341323f021e34173cb3824c
*/

// 跳跃表的实现
const (
	// DefaultMaxLevel TODO(suping.liu01) redis 设置这两个值是有什么深意吧？
	DefaultMaxLevel = 64            // Redis 中跳跃表的节点最大层数
	DefaultPFactor  = float64(0.25) // redis 中的 p 值（如果一个节点有有第1到第i层的索引，那么它在第i+1层有索引的概率为 p）
)

var (
	ErrNotFound = errors.New("SkipListNode not found")
)

type SkipListNode struct {
	Index uint64          // 索引值
	Val   interface{}     // 对象值, 真实存储的值
	Next  []*SkipListNode // 节点在不同层的下一个节点
}

func NewSkipListNode(index uint64, val interface{}, level int) *SkipListNode {
	return &SkipListNode{
		Index: index,
		Val:   val,
		Next:  make([]*SkipListNode, level),
	}
}

func (node *SkipListNode) toIndexNode() {
	// 只保留索引，不存储实际值的节点
	node.Val = nil
}

type SkipList struct {
	Head         *SkipListNode
	CurrentLevel int // 跳表当前的实际层数
}

func NewSkipList(head *SkipListNode, currentLevel int) *SkipList {
	return &SkipList{
		Head:         head,
		CurrentLevel: currentLevel,
	}
}

// Add 跳跃表添加节点
func (sl *SkipList) Add(node *SkipListNode) {
	if node == nil {
		return
	}

	// 随机索引层数
	level := randomLevel()
	node.Next = make([]*SkipListNode, randomLevel())

	searchNode := sl.Head // 游标
	// 从当前最高层开始, 逐层查询，添加索引，直到最底层
	for i := sl.CurrentLevel - 1; i >= 0; i-- {
		// 在每一层，从表头开始，遍历链表，直到找到 >= 待插入值的节点
		for searchNode.Next[i] != nil && searchNode.Next[i].Index < node.Index {
			searchNode = searchNode.Next[i]
		}

		// 如果当前正在遍历的层号 < 这个节点预期要创建索引的层数, 那么这一层就需要建立索引，即链表插入
		if i < level {
			if i != 0 {
				// 最底层才存储实际数据，>=0 层的都是索引，只保留索引，不记录数据(这样做是为了节省空间)
				node.toIndexNode()
			}
			// 把 node 插入到 searchNode 后面
			tmp := searchNode.Next[i]
			searchNode.Next[i] = node
			node.Next[i] = tmp
		}
	}

	// 如果这个节点预期要创建索引的层数 > 当前正在遍历测层数, 那当前层往上的层都要创建新的索引，而且在 head 之后
	if level > sl.CurrentLevel {
		node.toIndexNode()
		for i := sl.CurrentLevel; i < level; i++ {
			sl.Head.Next[i] = node
		}
	}

	// 更新当前层
	sl.CurrentLevel = level
}

func (sl *SkipList) Search(index uint64) (*SkipListNode, error) {
	var searchNode = sl.Head // 游标

	// 从当前最高层开始, 逐层查询
	for i := sl.CurrentLevel - 1; i >= 0; i-- {
		// 在每一层，从表头开始，遍历链表，直到找到 == 待插入值的节点
		for searchNode.Next[i] != nil && searchNode.Next[i].Index < index {
			searchNode = searchNode.Next[i]
		}
	}

	// 如果在当前层已经找到这个节点，那进入下一层继续判断，直到最底层获取到有值的节点
	if searchNode.Next[0] != nil && searchNode.Next[0].Index == index {
		return searchNode.Next[0], nil
	}

	return nil, ErrNotFound
}

func (sl *SkipList) Update(index uint64, Val interface{}) error {
	node, err := sl.Search(index)
	if err != nil {
		return err
	}

	node.Val = Val
	return nil
}

func (sl *SkipList) Remove(index uint64) {
	var searchNode = sl.Head // 游标

	// 从当前最高层开始, 逐层查询
	for i := sl.CurrentLevel - 1; i >= 0; i-- {
		// 在每一层，从表头开始，遍历链表，直到找到 == 待插入值的节点
		for searchNode.Next[i] != nil && searchNode.Next[i].Index < index {
			searchNode = searchNode.Next[i]
		}

		// 如果在当前层已经找到这个节点，删除这个节点并且向下层继续查询
		if searchNode.Next[i] != nil && searchNode.Next[i].Index == index {
			// 删除节点
			tmp := searchNode.Next[i].Next[i]
			searchNode.Next[i] = tmp
			// 进入下一层
			continue
		}
	}
	return
}

// ========== 工具 ==========
func randomLevel() int {
	level := 1 // 初始 level = 1
	for rand.Float64() < DefaultPFactor && level < DefaultMaxLevel {
		// 在不高于最大层数的前提, 有DEFAULT_P_FACTOR的概率在下一层也有索引
		level++
	}
	return level
}

