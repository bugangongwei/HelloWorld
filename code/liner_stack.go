package code

import (
	"fmt"
)

// [0,1,0,2,1,0,1,3,2,1,2,1]
func RianWater(arr []int) int {
	var sum = 0
	lmax := arr[0]
	for i := 0; i < len(arr); i++ {
		lmax = max(lmax, arr[i])

		right := i + 1
		rmax := arr[i]
		for right < len(arr) {
			if arr[right] > rmax {
				rmax = arr[right]
			}
			right++
		}

		if h := min(lmax, rmax) - arr[i]; h > 0 {
			sum += h
		}
		fmt.Println("current: ", arr[i], "lmax: ", lmax, "rmax: ", rmax)
	}

	return sum
}

func RianWater1(arr []int) int {
	// [i] 的水, min(lmax, rmax) - arr[i]
	sum := 0

	//  从左到右遍历, 得到所有的 lmax
	lmaxs := make([]int, len(arr))
	lmaxs[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		lmaxs[i] = max(arr[i], lmaxs[i-1])
	}

	// 从右到左遍历, 得到所有的 rmax
	rmaxs := make([]int, len(arr))
	rmaxs[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		rmaxs[i] = max(arr[i], rmaxs[i+1])
	}

	for i := 0; i < len(arr); i++ {
		h := min(lmaxs[i], rmaxs[i]) - arr[i]
		if h > 0 {
			sum += h
		}
	}

	return sum
}

// [0,1,0,2,1,0,1,3,2,1,2,1]
// 承接雨水的计算方式是 min(leftMax, rightMax) * (rightIdx - leftIdx - 1) - mid
// 栈顺序, 栈顶到栈底, 由小到大
// arr[i] > top(): 出现凹槽, 首先计算最近的凹槽, 即栈中前两个元素和当前元素组成的凹槽, for
// arr[i] < top(): 入栈
// arr[i] == top(): 出栈再入栈, 原因是, 我们希望用最右边的元素来作为凹槽的左边
func RianWater2(arr []int) int {
	var (
		// 元素下标的栈
		stack = new(StackInt)
		// 雨水
		sum = 0
	)

	stack.Push(0)
	for i := 1; i < len(arr); i++ {
		fmt.Println("stack: ", stack.len, arr[stack.Top()], "arr[i] ", arr[i])

		if arr[i] < arr[stack.Top()] {
			stack.Push(i)
		} else if arr[i] == arr[stack.Top()] {
			stack.Pop()
			stack.Push(i)
		} else {
			for stack.GetLen() > 0 && arr[i] > arr[stack.Top()] {
				// 右边 arr[i]
				// 中间 arr[stack.Pop()]
				// 左边 arr[stack.Top()]
				mid := stack.Pop()
				// 长
				if stack.GetLen() > 0 {
					fmt.Println("len from ", stack.Top(), " to ", i)
					len := i - stack.Top() - 1
					// 宽
					h := min(arr[stack.Top()], arr[i]) - arr[mid]
					// 凹槽 heights[mid]
					// 雨水
					if area := len * h; area > 0 {
						sum += area
					}
				}
			}
			stack.Push(i)
		}

	}

	return sum
}

// 2,1,5,6,2,3
func Juxing(arr []int) int {
	var res = 0
	for i := 0; i < len(arr); i++ {
		left, right := 0, 0
		for l := i - 1; l >= 0; l-- {
			if arr[l] >= arr[i] {
				left += arr[i]
			} else {
				break
			}
		}

		for r := i + 1; r < len(arr); r++ {
			if arr[r] >= arr[i] {
				right += arr[i]
			} else {
				break
			}
		}

		fmt.Println("left: ", left, "current: ", arr[i], "right: ", right)
		res = max(res, left+arr[i]+right)
	}

	return res
}

// 2,1,5,6,2,3
func Juxing1(heights []int) int {
	var (
		lidxs = make([]int, len(heights))
		ridxs = make([]int, len(heights))
		n     = len(heights)
		max   = 0
	)

	for i := 0; i < n; i++ {
		lidxs[i] = -1
		ridxs[i] = n
	}

	// 从左向右遍历, 得到所有 lidxs
	for i := 1; i < len(heights); i++ {
		if heights[i-1] < heights[i] {
			lidxs[i] = i - 1
		} else {
			for j := i - 1; j >= 0; j-- {
				if heights[j] < heights[i] {
					lidxs[i] = j
					break
				}
			}
		}
	}

	// 从右到左遍历, 得到所有 ridxs
	for i := n - 2; i >= 0; i-- {
		if heights[i+1] < heights[i] {
			ridxs[i] = i + 1
		} else {
			for j := i + 1; j < n; j++ {
				if heights[j] < heights[i] {
					ridxs[i] = j
					break
				}
			}
		}
	}
	for i := 0; i < len(heights); i++ {
		area := (ridxs[i] - lidxs[i] - 1) * heights[i]
		if area > max {
			max = area
		}
	}

	return max
}

// 2,1,5,6,2,3
// 和接雨水相同, 需要注意的是, 接雨水的题 首尾的不形成凹槽的元素可以直接舍弃
// 而本题中, 每一个元素都不能舍弃, 因为每一个元素都有可能创造出最大的矩形
func Juxing2(heights []int) int {
	var (
		// elem 就是数组元素下标
		stack = new(StackInt)
		res   = 0
	)

	stack.Push(0)
	for i := 1; i < len(heights); i++ {
		if heights[i] < heights[stack.Top()] {
			for stack.GetLen() > 0 && heights[i] < heights[stack.Top()] {
				mid := stack.Pop()
				if stack.GetLen() > 0 {
					res = max(res, (i-stack.Top()-1)*heights[mid])
				} else {
					res = max(res, (i-1)*heights[mid])
				}
			}
		}
		stack.Push(i)
	}

	// [2, 2, 4]
	for stack.GetLen() > 0 {
		mid := stack.Pop()
		if stack.GetLen() > 0 {
			res = max(res, (len(heights)-stack.Top()-1)*heights[mid])
		} else {
			res = max(res, (len(heights))*heights[mid])
		}
	}

	return res
}
