package code

import (
	"container/heap"
	_ "container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
双指针/滑动窗口(变长)
指针i: [0,len(s)-1), 遇到 s[j] num > 0, 则
指针j: [i+1, len(s)-1]

*/
func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var (
		i         = 0
		j         = 1
		maxResutl = 0
		m         = make(map[byte]struct{})
	)
	m[s[i]] = struct{}{}

	for j < len(s) && i < j {
		fmt.Println(m, i, j)
		if _, ok := m[s[j]]; ok {
			if j-i > maxResutl {
				fmt.Println(j - i)
				maxResutl = max(j-i, maxResutl)
			}

			for k := i; k < j; k++ {
				if s[j] == s[k] {
					i = k + 1
					break
				} else {
					delete(m, s[k])
				}
			}
		} else {
			if j == len(s)-1 {
				maxResutl = max(j-i+1, maxResutl)
			}
			m[s[j]] = struct{}{}
		}
		j++

		if i+maxResutl > len(s)-1 {
			break
		}
	}

	return maxResutl
}

/*
strs = ["flower","flow","flight"]
strs = ["dog","racecar","car"]
strs = ["","dog"]
*/
func LongestCommonPrefix(strs []string) string {
	var (
		minL        = len(strs[0])
		minStrIndex = 0
		comStr      = make([]rune, 0)
	)

	for i, str := range strs {
		if len(str) < minL {
			minL = len(str)
			minStrIndex = i
		}
	}

	for i, r := range strs[minStrIndex] {
		sum := 0
		for _, str := range strs {
			if []rune(str)[i] != r {
				break
			}
			sum++
		}
		if sum != len(strs) {
			break
		}
		comStr = append(comStr, r)
	}

	return string(comStr)
}

/*
s1 = "ab" s2 = "eidbaooo"
s1= "ab" s2 = "eidboaoo"
s1="abbcd",s2="eaiacbbodoo"

定长滑动窗口/双指针
窗口长度 = len(s1) = n
遍历顺序: 0<=i<len(s2)-n || 0+n<=j<len(s2)
窗口进: s2[j]
窗口出: s2[i]
其中, j-i = 2

窗口进: 字符的个数+1
窗口出: 字符的个数-1
最终: 在某个窗口, 每个字符的个数都和s1的字符个数相同, 则说明匹配
*/
func CheckInclusion(s1 string, s2 string) bool {
	var (
		// 记录 26 个小写字母, 每个字母的数量
		chaNumArray = [26]int{}
		// s2 的字符和 s1 的字符之间的 diff
		diff = 0
		n, m = len(s1), len(s2)
	)

	// 临界条件,省去不少计算
	if n > m {
		return false
	}

	// 初始化第一个窗口, len(s1) 在 s2 上圈中第一波值
	// 对于 s1, 应该往下减, 负数的绝对值表示 s1 中该字符的个数
	// 对于 s2, 应该往上增, 表示 s2 中该字符的个数
	for i, r := range s1 {
		chaNumArray[r-'a']--
		chaNumArray[s2[i]-'a']++
	}

	fmt.Println(chaNumArray)

	// 初始化 diff
	for _, num := range chaNumArray {
		if num != 0 {
			diff++
		}
	}

	// 第一个窗口的比对结果
	if diff == 0 {
		return true
	}

	// 第一个窗口不成功, 则从第二个窗口开始遍历
	for j := n; j < m; j++ {
		x, y := s2[j]-'a', s2[j-n]-'a'

		if x == y {
			continue
		}

		// 进入窗口的值, 0->1, diff++
		if chaNumArray[x] == 0 {
			diff++
		}
		chaNumArray[x]++
		// 添加数量之后, 查看值是否为 0, 为 0 表示这个字符平衡了
		// 进入窗口的值, 1->0, diff--
		if chaNumArray[x] == 0 {
			diff--
		}

		// 离开窗口的值, 0->-1, diff++
		if chaNumArray[y] == 0 {
			diff++
		}
		chaNumArray[y]--
		// 减去数量之后, 查看值是否为 0, 为 0 表示这个字符平衡了
		// 离开窗口的值, -1->0, diff--
		if chaNumArray[y] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}

	return false
}

/*
规律题
num1 = "123", num2 = "45"
“5535”
*/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var (
		m   = len(num1)
		n   = len(num2)
		arr = make([]int, m+n)
	)

	for i, r1 := range num1 {
		for j, r2 := range num2 {
			fmt.Println(int(r1-'0'), int(r2-'0'))
			num := int(r1-'0') * int(r2-'0')
			arr[i+j] += num / 10
			arr[i+j+1] += num % 10
			fmt.Println(arr)
		}
	}

	for i := m + n - 1; i > 0; i-- {
		if arr[i] >= 10 {
			fmt.Println(arr[i], arr[i]%10, arr[i]/10, arr[i-1])
			arr[i-1] = arr[i-1] + (arr[i] / 10)
			arr[i] = arr[i] % 10
		}
	}

	fmt.Println(arr)
	start := false
	var result string
	for i := 0; i < m+n; i++ {
		if arr[i] > 0 && !start {
			start = true
		}

		if start {
			fmt.Println(arr[i])
			result += strconv.Itoa(arr[i])
		}
	}

	return result
}

/*
s = "Alice does not even like bob"
"bob like even not does Alice"
s = "  Bob    Loves  Alice   "
"Alice Loves Bob"
*/
func ReverseWords(s string) string {
	prefix := " "
	subffix := " "
	split := "  "

	for strings.HasPrefix(s, prefix) {
		s = strings.TrimPrefix(s, prefix)
	}

	for strings.HasSuffix(s, subffix) {
		s = strings.TrimSuffix(s, subffix)
	}

	for strings.Contains(s, split) {
		s = strings.ReplaceAll(s, split, " ")
	}

	fmt.Println(s)
	words := strings.Split(s, " ")

	if len(words) == 0 {
		return ""
	}

	var (
		i = 0
		j = len(words) - 1
	)

	for i < j {
		// fmt.Println("before", words, i, j)
		revertSlice(words, i, j)
		// fmt.Println("after", words, i, j)
		i++
		j--
	}

	result := words[0]
	for k := 1; k < len(words); k++ {
		result += " " + words[k]
	}

	return result
}

func revertSlice(s []string, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func RestoreIpAddresses(s string) []string {
	var (
		// 单个 ip 地址
		path = make([]string, 0, 4)
		// 所有满足条件的 ip 地址
		result = make([]string, 0)
		// 递归函数
		fn func(startIdx int)
	)

	fn = func(startIdx int) {
		if startIdx > len(s)-1 {
			return
		}

		// 这种条件下的最后一个比较特殊, 必须是全部字符
		if len(path) == 3 {
			if len(s)-startIdx > 3 {
				// fmt.Println("最后一个不符合长度要求")
				return
			}

			str := s[startIdx:]
			if len(str) > 1 && strings.HasPrefix(str, "0") {
				return
			}
			// fmt.Println(startIdx, endIdx)
			num, _ := strconv.ParseInt(str, 10, 64)
			if num < 0 || num > 255 {
				return
			}

			path = append(path, str)
			result = append(result, strings.Join(path, "."))
			// fmt.Println("result: ", result)
			path = path[0 : len(path)-1]
			return
		}

		// if len(path) == 4 {
		// 	result = append(result, strings.Join(path, "."))
		// 	fmt.Println("result: ", result)
		// 	return
		// }

		for endIdx := startIdx; endIdx < len(s); endIdx++ {
			if endIdx-startIdx+1 > 3 {
				return
			}

			str := s[startIdx : endIdx+1]
			if len(str) > 1 && strings.HasPrefix(str, "0") {
				return
			}
			// fmt.Println(startIdx, endIdx)
			num, _ := strconv.ParseInt(str, 10, 64)
			if num < 0 || num > 255 {
				return
			}
			path = append(path, str)
			//fmt.Println("添加 path: ", path)
			// 递归
			fn(endIdx + 1)
			// 回溯
			path = path[0 : len(path)-1]
			//fmt.Println("回溯 path: ", path)
		}
	}

	fn(0)

	return result
}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
*/
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	var result = make([][]int, 0)

	// 1.排序, 从小到大
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 双指针
	// i 遍历
	// l 左指针, 代表第二个数
	// r 右指针, 代表第三个数
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return result
}

/*
图的深度优先遍历
递归写法
*/
func MaxAreaOfIsland(grid [][]int) int {
	var ans = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(ans, dfs(grid, i, j))
		}
	}

	return ans
}

func dfs(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		if grid[i][j] == 1 {
			grid[i][j] = 0
			sum := 1
			for _, dir := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				sum += dfs(grid, i+dir[0], j+dir[1])
			}
			return sum
		}
	}

	return 0
}

/*
旋转得到的部分有序数组
对中点来说, 左右两边总有一边是有序的
*/
func SearchDans(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var (
		left  = 0
		right = len(nums) - 1
	)

	for left <= right {
		mid := (right + left) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// 左边有序
			if target >= nums[left] && target < nums[mid] {
				// 值在左边
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右边有序
			if target > nums[mid] && target <= nums[right] {
				// 值在右边
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}

/*
nums = [1,3,5,4,7]
输出：3
*/
func FindLengthOfLCIS(nums []int) int {
	var (
		i   = 0
		ans = 0
		j   = 1
	)

	if len(nums) < 2 {
		return len(nums)
	}

	for ; j < len(nums); j++ {
		if nums[j] <= nums[j-1] {
			ans = max(ans, j-i)
			i = j
		}

		if j == len(nums)-1 {
			ans = max(ans, j-i+1)
		}
	}

	return ans
}

func FindKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})

	n := 1
	for _, v := range nums {
		if n == k {
			return v
		} else {
			n++
		}
	}

	return 0
}

/*
最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
[100,4,200,1,3,2]
[0,3,7,2,5,8,4,6,0,1]
*/
func LongestConsecutive(nums []int) int {
	var res = 0
	// 全部都丢到 map 中去
	var values = make(map[int]bool)
	for _, v := range nums {
		values[v] = true
	}

	// map 不为空的情况下, 遍历所有的值, 直到找到结果
	for i := 0; i < len(nums); i++ {
		var tmp int
		if values[nums[i]] {
			tmp++
			delete(values, nums[i])

			up := nums[i] + 1
			for len(values) > 0 && values[up] {
				tmp++
				delete(values, up)
				up++
			}

			d := nums[i] - 1
			for len(values) > 0 && values[d] {
				tmp++
				delete(values, d)
				d--
			}
		}
		res = max(res, tmp)
	}
	return res
}

/*
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。
*/
func GetPermutation(n int, k int) string {
	// 初始问题: 从有序的 [1,2...n] n 堆大小为 (n-1)! 的有序数组中, 找到第 k 个数
	// 首先定位 k 在 n 堆中属于第 seq 堆, 然后确定这个堆下标 (seq-1) 对应的元素 elem
	// 之后化解成子问题: 从有序的 [1,2...n]/a (n-1) 堆大小为 (n-2)! 的有序数组中, 找到第 k-(i-1)*(n-1)! 个数
	// 注意 arr 一致在变化, 找到一个 elem 就从 arr 中删除, 因为元素不可重复
	var (
		// 递归函数
		fn func(n, k, all int, arr []int)
		// 剩余可用来排列的元素
		arr = make([]int, 0, n)
		// 结果数组
		res string
	)

	// 剩余元素初始化
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	fmt.Println("我的初始数组对了吗? ", arr)

	fn = func(n, k, all int, arr []int) {
		if n < 2 {
			// 从 1 个元素里面找, 只能返回这个唯一的元素了
			res += strconv.Itoa(arr[0])
			return
		}

		fmt.Println("子问题分拆对了吗: ", n, k, all, arr)
		fmt.Println("我的结果对了吗: ", res)

		// 下一级阶乘(堆的大小)
		all = all / n
		// k 在第 seq 堆元素中
		seq := (k-1)/all + 1
		// 对应的元素
		elem := arr[seq-1]
		// 保存结果
		res += strconv.Itoa(elem)
		fmt.Println("我的元素对了吗? ", arr, seq-1, elem, strconv.Itoa(elem))
		// 删除已解决的 elem, 得到子问题的 arr
		var narr []int
		if seq == n {
			narr = arr[0 : seq-1]
		} else {
			narr = append(arr[0:seq-1], arr[seq:]...)
		}

		// 子问题, 新的元素个数, 新的 k, 新阶乘, 新的剩余元素
		fn(n-1, k-(seq-1)*all, all, narr)
	}

	fn(n, k, jc(n), arr)
	return res
}

func jc(n int) int {
	sum := 1
	for n > 0 {
		sum = sum * n
		n--
	}
	return sum
}

// 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// 输出：2
// 矩阵 n*n
func FindCircleNum(isConnected [][]int) int {
	// i-from, j-to
	// i, j 其实都是从 0 到 n-1
	var (
		// 城市个数
		n = len(isConnected)
		// 标志城市是否被访问过
		flags = make([]bool, n)
		// 省份数
		ans = 0
		// 递归函数
		dfs func(from int)
	)

	dfs = func(from int) {
		fmt.Println("开始: ", from)
		// 标记已访问
		flags[from] = true
		// 寻找 to, 为 1 的情况下才有相连的城市
		for j, to := range isConnected[from] {
			if to == 1 && !flags[j] {
				fmt.Println("相连的城市: ", j)
				// 从相连城市发起深度遍历
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		// 一开始所有的 from 都没被访问, 但是如果在 dfs 过程中被访问了, 那它就不需要重复去做 dfs 了
		if !flags[i] {
			ans++
			dfs(i)
		}
	}

	return ans
}

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println("我排序排对了吗? ", intervals)

	var (
		stack  = new(StackInt)
		result = make([][]int, 0)
	)

	for _, interval := range intervals {
		if stack.GetLen() == 0 {
			stack.Push(interval[0]) // start
			stack.Push(interval[1]) // end
		} else {
			if interval[0] <= stack.Top() {
				end := max(stack.Pop(), interval[1])
				start := min(stack.Pop(), interval[0])
				stack.Push(start)
				stack.Push(end)
			} else {
				// 前面一组出栈
				end := stack.Pop()
				start := stack.Pop()
				result = append(result, []int{start, end})
				// 后面一组入栈
				stack.Push(interval[0]) // start
				stack.Push(interval[1]) // end
			}
		}
	}

	end := stack.Pop()
	start := stack.Pop()
	result = append(result, []int{start, end})

	return result
}

/*
又是接雨水
双指针和动态规划的时间复杂度太高了
这里用单调栈来实现
以三个元素为一个凹槽[栈顶第二个元素, 栈顶元素, 当前元素]
(1) 单调栈从栈顶到栈底是单调递增的
(2)
if current > top: (min(left, right)-current)*(r-l-1), 水>0 才计入总数
if current == top: 入栈
if current < top: 入栈
(3) 单调栈的元素
下标
*/
func Trap(height []int) int {
	var (
		// 栈顶 <= 栈底的单调栈
		stack = new(StackInt)
		// 雨水总量
		sum = 0
	)
	stack.Push(0)

	for i := 1; i < len(height); i++ {
		if height[i] > height[stack.Top()] {
			for stack.GetLen() > 0 && height[i] > height[stack.Top()] {
				mid := stack.Pop()
				var left int
				if stack.GetLen() > 0 {
					left = stack.Top()
				} else {
					break
				}

				right := i
				h := min(height[left], height[right]) - height[mid]
				w := right - left - 1
				sum += h * w
			}
		}
		stack.Push(i)
	}

	return sum
}

/*
合并有序链表
双指针/穿针引线法
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		for l1 != nil && l2 != nil {
			fmt.Println("before: ", l1, l2)
			if l1.Val <= l2.Val {
				// 头指针
				if head.Next == nil {
					head.Next = l1
				}

				// 穿针引线
				tmp := l1.Next
				if tmp != nil && tmp.Val <= l2.Val {
					l1.Next = tmp
				} else {
					l1.Next = l2
				}

				// 指针后移
				l1 = tmp
			} else {
				// 头指针
				if head.Next == nil {
					head.Next = l2
				}

				// 穿针引线
				tmp := l2.Next
				if tmp != nil && tmp.Val < l1.Val {
					l2.Next = tmp
				} else {
					l2.Next = l1
				}

				// 指针后移
				l2 = tmp
			}
			fmt.Println("after: ", l1, l2)
		}
	}

	return head.Next
}

/*
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		// 进位
		tmp = 0
		// 新链表
		pHead = &ListNode{}
		pre   = pHead
	)

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 == nil && l2 != nil {
			sum = l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum = l1.Val
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}

		if tmp > 0 {
			sum += tmp
		}

		tmp = sum / 10

		current := &ListNode{Val: sum % 10}
		pre.Next = current

		pre = current
	}

	if tmp > 0 {
		pre.Next = &ListNode{Val: tmp % 10}
	}

	return pHead.Next
}

func MergeSort(s1, s2 *ListNode) *ListNode {
	var (
		p1, p2 = s1, s2
		head   *ListNode
	)

	if s1 == nil {
		return s2
	} else if s2 == nil {
		return s1
	} else {
		for p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				if head == nil {
					head = s1
				}

				// p1.Next 和 p2 选一个做为下一个数
				tmp := p1.Next
				if tmp != nil && tmp.Val < p2.Val {
					p1.Next = tmp
				} else {
					p1.Next = p2
				}
				p1 = tmp
			} else {
				if head == nil {
					head = s2
				}

				// p2.Next 和 p1 选一个做为下一个数
				tmp := p2.Next
				if tmp != nil && tmp.Val < p1.Val {
					p2.Next = tmp
				} else {
					p2.Next = p1
				}
				p2 = tmp
			}
		}
	}

	return head
}

type MinHeapListNode []*ListNode

func (mp *MinHeapListNode) Len() int {
	return len(*mp)
}

func (mp *MinHeapListNode) Less(i, j int) bool {
	return (*mp)[i].Val < (*mp)[j].Val
}

func (mp *MinHeapListNode) Swap(i, j int) {
	(*mp)[i], (*mp)[j] = (*mp)[j], (*mp)[i]
}

func (mp *MinHeapListNode) Push(v interface{}) {
	(*mp) = append((*mp), v.(*ListNode))
}

func (mp *MinHeapListNode) Pop() interface{} {
	p := (*mp)[len(*mp)-1]
	(*mp) = (*mp)[:len(*mp)-1]
	return p
}

func MergeKLists(lists []*ListNode) *ListNode {
	h := new(MinHeapListNode)
	for _, root := range lists {
		if root != nil {
			heap.Push(h, root)
		}

	}

	fmt.Println("链表个数: ", h.Len())

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}

		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}

type TmpData struct {
	idxI int
	idxJ int
	val  int
}

type MinHeapTmp []*TmpData

func (h *MinHeapTmp) Len() int           { return len(*h) }
func (h *MinHeapTmp) Less(i, j int) bool { return (*h)[i].val < (*h)[j].val }
func (h *MinHeapTmp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeapTmp) Push(v interface{}) { *h = append((*h), v.(*TmpData)) }
func (h *MinHeapTmp) Pop() interface{} {
	n := len(*h)
	p := (*h)[n-1]
	*h = (*h)[:n-1]
	return p
}

func KthSmallest(matrix [][]int, k int) int {
	if k <= 0 {
		return -1
	}

	mp := new(MinHeapTmp)
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > 0 {
			heap.Push(mp, &TmpData{idxI: i, idxJ: 0, val: matrix[i][0]})
		}
	}

	fmt.Println("堆大小", len(*mp))

	count := 0
	for mp.Len() > 0 {
		d := heap.Pop(mp).(*TmpData)
		fmt.Println("Pop: ", d)
		count++
		if count == k {
			return d.val
		}

		j := d.idxJ + 1
		if j < len(matrix[d.idxI]) {
			heap.Push(mp, &TmpData{val: matrix[d.idxI][j], idxI: d.idxI, idxJ: j})
		}
	}

	return -1
}

/*
输入：head = [4,2,1,3]
输出：[1,2,3,4]

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

输入：head = []
输出：[]
*/
func SortList(head *ListNode) *ListNode {
	return sort1(head, nil)
}

func sort1(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge1(sort1(head, mid), sort1(mid, tail))
}

func merge1(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2

	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}

	if tmp1 != nil {
		tmp = tmp1
	}

	if tmp2 != nil {
		tmp = tmp2
	}

	return dummyHead.Next
}

func MaximalSquare(matrix [][]byte) int {
	// dp[i][j] 下标定义: 以[i, j]元素为右下角元素的最大正方形的边长
	// 递推公式: dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])+1 (matrix[i][j]=1)
	// 递推公式: dp[i][j] = 0 (matrix[i][j]=0)
	// 初始值: dp[i][0] = 0 (matrix[i][0]=0), dp[i][0] = 1 (matrix[i][0]=1)
	// 初始值: dp[0][j] = 0 (matrix[0][j]=0), dp[0][j] = 1 (matrix[0][j]=1)
	// 遍历方式, 从martrix[1][1] 开始遍历
	// 举例验证: matrix 为空, return 0

	if len(matrix) == 0 {
		return 0
	}

	var (
		m, n = len(matrix), len(matrix[0])
		dp   = make([][]int, m)
		max  = 0
	)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			max = Max(max, 1)
			dp[i][0] = 1
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			max = Max(max, 1)
			dp[0][j] = 1
		}
	}

	fmt.Println("看看我的初始值 ", dp, max)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j] == '1' && matrix[i][j-1] == '1' && matrix[i-1][j-1] == '1' {
					dp[i][j] = Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				} else {
					dp[i][j] = 1
				}

				max = Max(max, dp[i][j])
			}
		}
	}

	return max * max
}

func Max(params ...int) int {
	max := params[0]

	for _, p := range params {
		if p > max {
			max = p
		}
	}

	return max
}

func Min(params ...int) int {
	min := params[0]

	for _, p := range params {
		if p < min {
			min = p
		}
	}

	return min
}

func ValidUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}

	count := bytesFollowed(data[0])
	if count == -1 {
		return false
	}

	fmt.Println("bytes followd: ", count)
	if len(data) < count+1 {
		return false
	}
	
	for i := 1; i < count+1; i++ {
		fmt.Println("?", data[i], data[i]>>6 != 2)
		if data[i]>>6 != 2 {
			return false
		}
	}

	if count+1 < len(data) && len(data[count+1:]) > 0 {
		return ValidUtf8(data[count+1:])
	}

	return true
}

func bytesFollowed(first int) int {
	// 0*******
	if first>>7 == 0 {
		return 0
	}

	// 110*****
	if first>>5 == 6 {
		return 1
	}

	// 1110****
	if first>>4 == 14 {
		return 2
	}

	// 11110***
	if first>>3 == 30 {
		return 3
	}

	return -1
}
