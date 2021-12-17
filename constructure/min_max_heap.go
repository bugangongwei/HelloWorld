package constructure

import (
	_ "container/heap"
	"fmt"
	"math"
)

// ======== 自行实现 ========
type MinHeap struct {
	Elements []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Elements: []int{math.MinInt64}}
}

func (mp *MinHeap) Push(v int) {
	mp.Elements = append(mp.Elements, v)

	i := len(mp.Elements) - 1
	for ; mp.Elements[i/2] > v; i = i / 2 {
		mp.Elements[i] = mp.Elements[i/2]
	}

	mp.Elements[i] = v
	fmt.Println(v, "插入在位置", i)
}

func (mp *MinHeap) Pop() int {
	n := len(mp.Elements)
	minElement := mp.Elements[1]
	lastElement := mp.Elements[n-1]

	fmt.Println("最小值", minElement)
	fmt.Println("最后值", lastElement)

	var child int
	for i := 1; i*2 < n; i = child {
		// child 表示左右孩子中比较小的那个的下标
		child = i * 2
		if child+1 < n && mp.Elements[child+1] < mp.Elements[child] {
			child++
		}

		fmt.Println("i是", i, "child 是", child, "值是 ", mp.Elements[child])

		if mp.Elements[child] < lastElement {
			mp.Elements[i] = mp.Elements[child]
		} else {
			break
		}
	}

	mp.Elements[child] = lastElement
	fmt.Println("child", child, "值", lastElement)
	mp.Elements = mp.Elements[0 : n-1]
	return minElement
}

func (mp *MinHeap) Top() int {
	return mp.Elements[1]
}

// ======== 使用 container/heap 包实现 ========
type MaxHeap []int

func (mp *MaxHeap) Len() int {
	return len(*mp)
}

func (mp *MaxHeap) Less(i, j int) bool {
	return (*mp)[i] > (*mp)[j]
}

func (mp *MaxHeap) Swap(i, j int) {
	(*mp)[i], (*mp)[j] = (*mp)[j], (*mp)[i]
}

func (mp *MaxHeap) Push(x interface{}) {
	*mp = append(*mp, x.(int))
}

func (mp *MaxHeap) Pop() interface{} {
	result := (*mp)[len(*mp)-1]
	*mp = (*mp)[:len(*mp)-1]
	return result
}
