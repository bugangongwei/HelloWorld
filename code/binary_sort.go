package code

import (
	"fmt"
)

/*
4. 寻找两个正序数组的中位数
A：1 2 3 6
B: 1 3 4 5 9 10

n = 4, m = 6, 中位数是第 (10/2=5) 和第 (10/2+1=6) 个数的平均值

题目变体： 在 log 的复杂度下找到第 k(k=5) 大的数
简单点：数组 A 中找 k/2 个， 数组 B 中找 k/2 个， k/2*2=k 就正好是第 k 个数
由于 k/2 不可能恰好是整数，当它是小数的时候， 进行向下取整， 然后 k/2*2 <= k 了，那 A 中 k/2 个数和 B 中 k/2 个数合起来之后，最大的那个数可能是第 k 个数或者小于第 k 个数；
这个不太好确定， 但是我们能够确定的是， 比较小的那 k/2 个数里面肯定不存在第 k 个数；

二分法的精髓：排除不可能包含查找结果的那一半数据
在这里， 就是把比较小的前 k/2 个数去掉

(1) 初始化 k = (m+n)/2 = 5
(2) 初始化 i, j = 0, 0
(3) （i, j 分别在 AB 的下标范围内循环）判断 A(i) 和 B(j)
if k == 1, 可以直接返回 (A(i) 和 B(j)) 中最小的一个;
if A(i) < B(j), 就把 A 中的 k/2 个数去掉，m = m - k/2 = 2, k = (m+n)/2 = (8/2) = 4，i = i + k/2 ；
否则，就把 B 中的 k/2 个数去掉，n = n - k/2, k = (m+n)/2，j = j + k/2 ；
(4) 如果找到最后，其中一个数组不够长，那么第 k 大的数就要在剩下的一个数组中找
如果是 i>m, 也就是数组 A 不够长，就在 B 数组中找到第 k 个数， 也就是 B(j-k/2+1+k-1) （前提：j+k-1 < n）
如果是 j>n, 也就是数组 B 不够长，就在 A 数组中找到第 k 个数， 也就是 A(i-k/2+1       +k-1) （前提：i+k-1 < m）

回归题目本身，不是找第 k 大的数，而是找中位数
对于奇数，第 k 大的数就是中位数
对于偶数，中位数是第 k 大的数和 k + 1 大的数的平均数
*/

// FindMedianSortedArrays 使用 log(m+n) 的时间复杂度来计算两个有序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		m, n = len(nums1), len(nums2)
		a    = m + n
	)

	if a&1 == 0 {
		// 偶数, 第 (a/2) 和第 (a/2)+1 大的数的平均值就是中位数
		x := findKthNumber(nums1, nums2, a/2)
		y := findKthNumber(nums1, nums2, a/2+1)
		fmt.Println(x, y)
		return float64(x+y) / 2
	}

	// 奇数， 第 (a+1) 个数就是中位数
	return float64(findKthNumber(nums1, nums2, (a+1)/2))
}

func findKthNumber(nums1 []int, nums2 []int, k int) int {
	// 初始化参数
	var (
		m, n = len(nums1), len(nums2)
		i, j = 0, 0
	)

	// 循环二分
	for i >= 0 && i < m && j >= 0 && j < n {
		// 得到一类条件下的结果
		if k == 1 && i < m && j < n {
			return min(nums1[i], nums2[j])
		}

		if i+k/2-1 < m && j+k/2-1 < n {
			if nums1[i+k/2-1] < nums2[j+k/2-1] {
				// 数组 nums1 去掉 k/2 个数
				i += k / 2
			} else {
				// 数组 nums2 去掉 k/2 个数
				j += k / 2
			}
			k -= k / 2
		} else if i+k/2-1 >= m {
			k -= m - i
			if nums1[m-1] < nums2[j+k/2-1] {
				// 数组 nums1 去掉 k/2 个数
				i = m
			} else {
				j += m - i
			}
		} else {
			k = k - (n - j)
			if nums1[i+k/2-1] < nums2[n-1] {
				j += n - j
			} else {
				j = n
			}
		}
	}

	// 不满足循环的情况下的单数组处理
	if i >= m && j+k-1 < n {
		// 数组 nums1 不够长
		return nums2[j+k-1]
	}

	if j >= n && i+k-1 < m {
		// 数组 nums2 不够长
		return nums1[i+k-1]
	}

	return -1
}
