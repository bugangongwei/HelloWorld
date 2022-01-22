package code

import (
	"log"
)

// 2059. 转化数字的最小运算数

// MinimumOperations 单向 BFS 遍历寻找最短路径
func MinimumOperations(nums []int, start int, goal int) int {
	// 定义广度优先遍历需要的 queue
	q := NewQueue()
	// 定义 map 来标记被访问过的节点, 在第几次转换时被访问
	m := make(map[int]int, 0)

	q.Push(start)
	m[start] = 0

	// 基于 queue 的广度优先遍历
	for q.Len() > 0 {
		curr := q.Pop()
		currStep := m[curr]

		for _, num := range nums {
			nexts := []int{curr + num, curr - num, curr ^ num}
			for _, next := range nexts {
				if next == goal {
					return currStep + 1
				}

				if next < 0 || next > 1000 {
					continue
				}

				if _, ok := m[next]; ok {
					continue
				}

				q.Push(next)
				m[next] = currStep + 1
			}
		}
	}

	return -1
}

// MinimumOperationsDouble 双向 BFS 遍历寻找最短路径
func MinimumOperationsDouble(nums []int, start int, goal int) int {
	// 双向维护 BFS
	q1, q2 := NewQueue(), NewQueue()
	m1, m2 := map[int]int{}, map[int]int{}

	q1.Push(start)
	q2.Push(goal)

	m1[start] = 0
	m2[goal] = 0

	for q1.Len() > 0 && q2.Len() > 0 {
		// 正反双向，平衡节点个数, 对入队数少的一个方向，尝试放入节点
		if q1.Len() > q2.Len() {
			q1, q2 = q2, q1
			m1, m2 = m2, m1
		}

		n := q1.Len()
		for i := 0; i < n; i++ {
			// 获取当前的节点和节点路径距离
			curr := q1.Pop()
			currStep := m1[curr]

			for _, num := range nums {
				nextSet := NewIntSet([]int{curr + num, curr - num, curr ^ num}...)
				for _, next := range nextSet.Keys() {
					// 一旦遇到对方向已经访问过的节点，把两边的长度相加返回
					if _, ok := m2[next]; ok {
						return m2[next] + currStep + 1
					}

					// 不符合条件或者已经访问过的节点，不再计算
					if _, ok := m1[next]; ok || next < 0 || next > 1000 {
						continue
					}

					// 节点加入队列
					q1.Push(next)
					m1[next] = currStep + 1
				}
			}
		}
	}

	return -1
}

type IntSet struct {
	Elems map[int]struct{}
}

func NewIntSet(elems ...int) *IntSet {
	var intSet = &IntSet{Elems: make(map[int]struct{})}
	for _, elem := range elems {
		intSet.Set(elem)
	}
	return intSet
}

func (set *IntSet) Set(elem int) {
	set.Elems[elem] = struct{}{}
}

func (set *IntSet) Keys() []int {
	var keys = make([]int, 0)
	for k, _ := range set.Elems {
		keys = append(keys, k)
	}
	return keys
}

type BFSQueue struct {
	Elems []int
}

func NewQueue() *BFSQueue {
	return &BFSQueue{
		Elems: make([]int, 0),
	}
}

func (q *BFSQueue) Len() int {
	return len(q.Elems)
}

func (q *BFSQueue) Push(elem int) {
	q.Elems = append(q.Elems, elem)
}

func (q *BFSQueue) Pop() int {
	if len(q.Elems) == 0 {
		log.Fatal("pop from empty queue")
	}

	result := q.Elems[0]
	q.Elems = q.Elems[1:]
	return result
}

// 127. 单词接龙

func LadderLength(beginWord string, endWord string, wordList []string) int {
	// 队列：待考察的节点
	q := []string{beginWord}
	// 有相同的两个字母的单词, 如 ["*og"] -> {"cog", "dog"}
	matches := map[string][]string{}
	// 节点与 beginWord 间的最短路径长度
	shortestPathLen := map[string]int{beginWord: 0}

	// 初始化 matches
	// ["*og"] -> {"cog", "dog"}
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			tmp := []rune(word)
			tmp[i] = '*'
			t := string(tmp)
			matches[t] = append(matches[t], word)
		}
	}

	for len(q) > 0 {
		for _, curr := range q {
			// 如果当前的单词已经是目标单词，直接返回最短路径
			if curr == endWord {
				// 注意：题目要的不是跳数，而是最短路径中有几个节点
				return shortestPathLen[curr] + 1
			}

			// 分别匹配不同的单词
			for i := 0; i < len(curr); i++ {
				tmp := []rune(curr)
				tmp[i] = '*'
				t := string(tmp)
				if _, ok1 := matches[t]; ok1 {
					for _, w := range matches[t] {
						// 已经加入过最短路径的单词不需要再加入
						if _, ok2 := shortestPathLen[w]; ok2 {
							continue
						}
						// 加入待考察
						q = append(q, w)
						// 最短路径 +1
						shortestPathLen[w] = shortestPathLen[curr] + 1
					}
					// 删除已经匹配的单词，以免后面形成环路
					delete(matches, t)
				}
			}

			q = q[1:]
		}
	}

	return 0
}

// 1345. 跳跃游戏 IV

func MinJumps(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// 数字所在的所有下标位置
	m := make(map[int][]int)
	// BFS 辅助需要的队列
	q := []int{0}
	// 跳数
	Deps := map[int]int{0: 0}

	// 构造 m
	for idx, val := range arr {
		m[val] = append(m[val], idx)
	}

	for len(q) > 0 {
		for _, currIdx := range q {
			// 当前考察的节点，和节点对应的跳数
			currDep := Deps[currIdx]

			// 如果当前考察的节点已经是最终节点了，直接返回跳数
			if currIdx == len(arr)-1 {
				return currDep
			}

			// 当前节点可以跳到的下一个节点的集合
			var nexts []int
			if v, ok := m[arr[currIdx]]; ok {
				nexts = v
				delete(m, arr[currIdx])
			}
			if _, ok1 := Deps[currIdx-1]; currIdx > 0 && !ok1 {
				nexts = append(nexts, currIdx-1)
			}
			if _, ok2 := Deps[currIdx+1]; currIdx < len(arr)-1 && !ok2 {
				nexts = append(nexts, currIdx+1)
			}

			// 把下一个目标作为待考察对象放入队列中
			for _, next := range nexts {
				if _, ok := Deps[next]; ok {
					continue
				}

				q = append(q, next)
				Deps[next] = currDep + 1
			}

			// 去掉已经考察过的对象
			q = q[1:]
		}
	}

	return 0
}
