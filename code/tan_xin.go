package code

import (
	"fmt"
	"sort"
	"strconv"
)

// 分发小饼干
// g: 小学生
// s: 小饼干
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var (
		// 小学生下标
		index int
		// result
		sum int
	)

	for i := 0; index < len(g) && i < len(s); i++ {
		if s[i] >= g[index] {
			sum++
			index++
		}
	}

	return sum
}

/*
摆动序列
[1,1,2,2,3,4,10,3,2,1,0,7,5,7,10,11,12]
*/
func WiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var (
		pre int
		cur int
		sum = 1
	)

	for i := 1; i < len(nums); i++ {
		cur = nums[i] - nums[i-1]
		if (pre >= 0 && cur < 0) || (pre <= 0 && cur > 0) {
			sum++
			pre = cur
		}
	}

	return sum
}

/*
买卖股票 贪心
[7,1,5,3,6,4]
*/
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		profit int
		sum    int
	)

	for i := 1; i < len(prices); i++ {
		profit = prices[i] - prices[i-1]
		if profit > 0 {
			sum += profit
		}
	}

	return sum
}

/*
加油站
gas = [2,3,4] cost = [3,4,3]
gas = [1,2,3,4,5] cost = [3,4,5,1,2]
*/
func CanCompleteCircuit(gas []int, cost []int) int {
	var (
		// 当前一段路的剩余油量, 如果出现剩余量 < 0的情况, 说明这条路走不通
		tmpRest = 0
		// 走完一圈的剩余油量, 如果最后剩余量 < 0, 说明这圈走不下来
		restSum = 0
		// result
		start = 0
	)

	for i := 0; i < len(gas); i++ {
		tmpRest += gas[i] - cost[i]
		restSum += gas[i] - cost[i]
		if tmpRest < 0 {
			// 尝试失败, 换一个 start
			tmpRest = 0
			start = i + 1
		}
	}

	if restSum < 0 {
		return -1
	}
	return start
}

/*
135. 分发糖果
[1, 2, 2, 5, 4, 3, 2]
*/
// 双向遍历
func Candy1(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	var (
		right int
		l2r   = make([]int, len(ratings))
		sum   = 0
	)

	// 从左到右, 右边的比左边的大, 就多分一个糖果, 否则, 就等于1
	l2r[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			l2r[i] = l2r[i-1] + 1
		} else {
			l2r[i] = 1
		}
	}

	// 从右到左, 左边的比右边的大, 则多分一个糖果, 否则就等于1
	right = max(1, l2r[len(ratings)-1])
	sum += right
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// 两个结果选最大的那个
			right = max(right+1, l2r[i])
		} else {
			right = max(1, l2r[i])
		}
		sum += right
	}

	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 单向遍历
// [1, 2, 2, 5, 4, 3, 2]
func Candy2(ratings []int) int {
	var (
		pre  = 1
		sum  = pre
		desc = 1
		inc  = 1
	)

	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			desc = 1
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			sum += pre
			inc = pre
		} else {
			if desc == inc {
				desc++
			}
			pre = 1
			sum += desc
			desc++
		}
	}

	return sum
}

/*
860.柠檬水找零
示例 2： 输入：[5,5,10] 输出：true
示例 3： 输入：[10,10] 输出：false
示例 4： 输入：[5,5,10,10,20] 输出：false
*/
func LemonadeChange(bills []int) bool {
	var (
		// 两个零钱抽屉
		five = 0
		ten  = 0
	)

	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if ten == 0 {
				if five < 3 {
					return false
				}
				five = five - 3
			} else {
				ten--
				if five == 0 {
					return false
				}
				five--
			}
		}
	}

	return true
}

/*
406. 根据身高重建队列
输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
*/

func ReconstructQueue(people [][]int) [][]int {
	// [[7,0],[7,1],[6,1],[5,0],[5,2],[4,4]]
	// [[7,0],[6,1],[7,1],[5,0],[5,2],[4,4]]
	// [[5,0],[7,0],[6,1],[7,1],[5,2],[4,4]]
	// [[5,0],[7,0],[5,2],[6,1],[7,1],[4,4]]
	// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

	// 排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// 贪心
	for i := 0; i < len(people); i++ {
		if i != people[i][1] {
			// 插入到 i 位置
			tmp := people[i]
			copy(people[tmp[1]+1:i+1], people[tmp[1]:i])
			people[tmp[1]] = tmp
		}
	}

	return people
}

/*
452. 用最少数量的箭引爆气球
[[10,16],[2,8],[1,6],[7,12]]
*/
func FindMinArrowShots(points [][]int) int {
	// 边界条件
	if len(points) == 0 {
		return 0
	}

	// 排序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	fmt.Println("排序后 ", points)

	var (
		// 可射击范围
		shot = points[0]
		// 总射击次数
		sum = 1
	)
	// 遍历引爆
	for i := 1; i < len(points); i++ {
		// fmt.Println("可射击范围 ", shot)
		// fmt.Println("sum ", sum)
		if points[i][0] <= shot[1] {
			// 折叠, 缩小可射击范围
			shot[0] = max(points[i][0], shot[0])
			shot[1] = min(points[i][1], shot[1])
		} else {
			// 不重叠
			sum++
			// 重置可射击范围
			shot = points[i]
		}
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
435. 无重叠区间
跟引爆气球不同的是, 这里相等不算重叠
逻辑一样
返回值变成(总数-非重叠数)
*/
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var (
		// 重叠区域
		shot = intervals[0]
		// 总重叠数
		sum = 1
	)

	// 贪心最优解
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < shot[1] {
			// 重叠, 缩小重叠范围
			shot[0] = max(shot[0], intervals[i][0])
			shot[1] = min(shot[1], intervals[i][1])
		} else {
			// 不重叠
			sum++
			shot = intervals[i]
		}
	}
	return len(intervals) - sum
}

/*
763.划分字母区间
S = "ababcbacadefegdehijhklij
*/
func PartitionLabels(s string) []int {
	var (
		runes    = []rune(s)
		farIndex = make(map[rune]int)
	)

	// 找到每个元素的最远出现位置
	for i, r := range runes {
		farIndex[r] = i
	}

	// 分割字符串
	var (
		maxIndex = 0
		result   = make([]int, 0)
		len      = 0
	)
	for i, r := range runes {
		len++
		if farIndex[r] > maxIndex {
			maxIndex = farIndex[r]
		}
		if i == maxIndex {
			result = append(result, len)
			len = 0
			maxIndex = 0
		}
	}

	return result
}

/*
56. 合并区间
intervals = [[1,3],[2,6],[8,10],[15,18]] 输出: [[1,6],[8,10],[15,18]]
*/
func Merge(intervals [][]int) [][]int {
	// 边界条件
	if len(intervals) == 0 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	// fmt.Println("排序后 ", intervals)

	// 贪心重叠
	var (
		// 重叠范围, 用来判断重叠
		overlap = intervals[0]
		// 合并范围, 用来最终返回
		merge = make([][]int, 0)
	)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= overlap[1] {
			// 重叠, 扩大重叠范围
			overlap[0] = min(overlap[0], intervals[i][0])
			overlap[1] = max(overlap[1], intervals[i][1])
		} else {
			// 不重叠, 把 merge 的 overlap 加入 merge 数组
			// fmt.Println("上一个 overlap ", overlap)
			merge = append(merge, overlap)
			overlap = intervals[i]
		}
	}
	merge = append(merge, overlap)
	return merge
}

/*
738.单调递增的数字
最大的各位递增的整数
示例 1: 输入: N = 10 输出: 9
示例 2: 输入: N = 1234 输出: 1234
示例 3: 输入: N = 332 输出: 299
*/
func MonotoneIncreasingDigits(N int) int {
	runes := []rune(strconv.Itoa(N))
	later := runes[len(runes)-1]

	for i := len(runes) - 2; i >= 0; i-- {
		fmt.Println(runes[i], later)
		if runes[i] > later {
			// 非递增
			for j := i + 1; j < len(runes); j++ {
				runes[j] = '9'
			}
			runes[i]--
		}
		later = runes[i]
	}

	res, _ := strconv.Atoi(string(runes))
	return res
}

/*
714. 买卖股票的最佳时机含手续费
prices = [1, 3, 2, 8, 4, 9], fee = 2 输出: 8
局部最优: 局部买卖区间内, 买入是在价格最低时, 卖出是在不再获取利润时, 则在买入和卖出期间, 每一个点都是在盈利的
全局最有: 全局买卖期间, 所有持有的点, 都是盈利的, 只要每个盈利点都计算进来, 最终的结果一定时最大的盈利数字
*/
func MaxProfitII(prices []int, fee int) int {
	var (
		result = 0
		// 输入的 prices 不会是空数组
		minPrice = prices[0]
	)

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] > minPrice+fee {
			result += prices[i] - minPrice - fee
			minPrice = prices[i] - fee
		}
	}

	return result
}

/*
968.监控二叉树
0: 无覆盖
1: 有摄像头
2: 有覆盖
*/
func MinCameraCover(root *TreeNode) int {
	var (
		result = 0
		fn     func(root *TreeNode) int
	)

	fn = func(root *TreeNode) int {
		if root == nil {
			return 2
		}

		l := fn(root.Left)
		r := fn(root.Right)

		if l == 0 || r == 0 {
			result++
			return 1
		}

		if l == 1 || r == 1 {
			return 2
		}

		return 0

	}

	if fn(root) == 0 {
		result++
	}
	return result
}
