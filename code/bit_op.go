package code

import (
	"math"
	"sort"
)

/*
https://leetcode-cn.com/problems/divide-two-integers/solution/acmjin-pai-ti-jie-bei-zeng-jia-bu-yong-l-mg8s/
这道题的关键点是: 根据题意, 32位的有符号正整数, 其实暗含一个整型溢出的条件
解体思想: 除法的本质是减法
优化的关键: 减少减法的次数
由循环dividend/divisor得到结果, 再优化成贪心的方式得到结果
*/
func Divide(dividend int, divisor int) int {
	// (1) 溢出情况, -2^31, 2^31-1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// (2) 除数为 0
	if divisor == 0 {
		return 0
	}

	// (3) 符号计算
	rev := false
	if dividend > 0 {
		rev = !rev
	} else {
		dividend = -dividend
	}

	if divisor > 0 {
		rev = !rev
	} else {
		divisor = -divisor
	}

	var res = 0
	for dividend >= divisor {
		sum, divided := divisor, 1
		for dividend >= (sum << 1) {
			sum = sum << 1
			divided = divided << 1
		}
		dividend -= sum
		res += divided
	}

	if rev {
		return -res
	}

	return res
}

/*
https://leetcode-cn.com/problems/subsets/solution/zi-ji-by-leetcode-solution/
回溯算法：
套用回溯算法的公式, 但是这里要注意, path每添加一个元素都要算做一种组合, 即每一层都有合理的结果
*/
func Subsets(nums []int) (ans [][]int) {
	var (
		path   = make([]int, 0)
		result = make([][]int, 0)
		fn     func(start int)
	)
	result = append(result, path)

	fn = func(start int) {
		if start >= len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			// 处理节点
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)

			// 递归
			fn(i + 1)

			// 回溯, 撤销处理结果
			path = path[0 : len(path)-1]
		}
	}

	fn(0)

	return result
}

/*
子集进阶
通过树图可以总结得出, 同一层的数据不允许有重复, 如果有重复, 那么会导致结果集有重复
*/
func SubsetsWithDup(nums []int) (ans [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	var (
		path = make([]int, 0)
		fn   func(idx int)
	)
	ans = append(ans, path)

	fn = func(idx int) {
		if idx >= len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			// 判断是否同层元素重复
			if i > 0 && nums[i] == nums[i-1] && i > idx {
				continue
			}

			// 添加结果集(注意复制和指针问题)
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)

			// fmt.Println("idx: ", idx, "i: ", i, "path: ", path, "ans: ", ans)
			// 递归到下一层
			fn(i + 1)

			// 回溯
			path = path[0 : len(path)-1]
		}
	}

	fn(0)

	return
}

/*
位运算的方式, 用01来表示元素是否存在
注意:解集不包含重复的子集, nums中没有重复元素, 如果有重复元素, 则需要去重
{5, 2, 9}
0 000 {}
1 001 {9}
2 010 {2}
3 011 {2, 9}
4 100 {5}
5 101 {5, 9}
6 110 {5, 2}
7 111 {5, 2, 9}

for i:=0;i<(1<<len(nums)-1);i++ {
	// i = 2 {010}
	i>>0&1=0
	i>>1&1=1 path.append(nums[1])
	i>>2&1=0
}

时间复杂度: O(1<<n * n)
空间复杂度: O(n)
*/
func Subsets2(nums []int) (ans [][]int) {
	for i := 0; i < 1<<len(nums); i++ {
		path := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				path = append(path, nums[j])
			}
		}
		ans = append(ans, path)
	}
	return ans
}

/*
子集进阶: 对可能出现重复的情况在迭代中及时避免掉
排序后, 对于相同的两个数: 01, 10 是一样的数据, 所以遇到 01 就不加入结果集, 00和11都是有意义的
*/
func SubsetsWithDup2(nums []int) (ans [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
outer:
	for i := 0; i < 1<<len(nums); i++ { // {000, 001, 010, 011, 101, 110, 111}
		path := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				if j > 0 && i>>(j-1)&1 == 0 && nums[j] == nums[j-1] {
					// xx01, 数值相等情况下, 默认和xx10的组合结果是一样的
					continue outer
				}
				path = append(path, nums[j])
			}
		}
		ans = append(ans, path)
	}
	return
}

/*
https://leetcode-cn.com/problems/single-number-ii/solution/zhi-chu-xian-yi-ci-de-shu-zi-ii-by-leetc-23t6/
O(n+logn)的时间复杂度
*/
func SingleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i+1 >= len(nums) || nums[i] < nums[i+1] {
			if i == 0 || nums[i] > nums[i-1] {
				return nums[i]
			}
		}
	}

	return 0
}

/*
使用位计算
规律: {1,2,2,2} -> {0001, 0010, 0010, 0010} -> {0001 + 0010 + 0010 + 0010} -> {0031}%3 -> {0001}
时间复杂度:O(32*n), 符合线性时间复杂度的要求
空间复杂度O(1)
*/
func SingleNumber2(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		var total int32
		for j := 0; j < len(nums); j++ {
			total += int32(nums[j]) >> i & 1
		}

		res += total % 3 << i
	}
	return int(res)
}
