package code

import "fmt"

/*
斐波那契数列
*/
func Fib(n int) int {
	// 1. 确定 dp 数组以及下标含义
	// dp[i] = 第 i 个斐波那契值

	// 2. 确定递推公式
	// dp[i] = dp[i-1] + dp[i-2]

	// 3. dp 数组初始化
	// dp[0] = 0
	// dp[1] = 1

	// 4. 确定遍历顺序
	// 从 0..n

	// 5. 举例推导 dp 数组
	// n = 5
	// 0,1,1,2,3,5
	if n == 0 {
		return 0
	}

	var (
		dp     = []int{0, 1}
		result = 0
	)

	for i := 2; i <= n; i++ {
		result = dp[0] + dp[1]
		fmt.Println(result)
		dp[0] = dp[1]
		dp[1] = result
	}

	return dp[1]
}

/*
70. 爬楼梯
*/
func ClimbStairs(n int) int {
	// 1.确定 dp 数组以及下标含义
	// dp[i]:走到第 i 个可以有 dp[i] 中方式

	// 2.确定递推公式
	// dp[i] = dp[i-1] + dp[i-2]

	// 3.dp 数组如何初始化
	// dp[0] = 0
	// dp[1] = 1

	// 4.确定遍历顺序
	// 0..n

	// 5.举例推导 dp 数组
	// dp = {1,1,2,3,5}
	if n == 0 {
		return 1
	}

	var (
		dp     = []int{1, 1}
		result = 0
	)

	for i := 2; i <= n; i++ {
		result = dp[0] + dp[1]
		fmt.Println(result)
		dp[0] = dp[1]
		dp[1] = result
	}

	return dp[1]
}

/*使用最小花费爬楼梯*/
func MinCostClimbingStairs(cost []int) int {
	// 1. dp[i] 下标的含义
	// 到达下标 i 的台阶所花费的最少体力
	// 2. 推导 dp[i] 的表达式
	// dp[i] = dp[i-1] + cost[i], 从 i-1 走一步上来
	// dp[i] = dp[i-2] + cost[i], 从 i-2 跨两步上来
	// dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	// 3. dp[i] 的初始化
	// dp[0] = cost[0]
	// dp[1] = cost[1]
	// 4. 确定遍历顺序
	// 2..len(cost)-1
	// 5. 举例推导 dp 数组
	// [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
	// dp[0] = 1
	// dp[1] = 100
	// dp[2] = min(1, 100)+1=2
	// dp[3] = min(100,2)+1=3
	// dp[4] = min(2,3)+1=3
	// dp[5] = min(3,3)+100 = 103
	// dp[6] = min(3,103)+1=4
	// dp[7] = min(103,4)+1=5
	// dp[8] = min(4,5)+100=104
	// dp[9] = min(5,104)+1=6

	var (
		dp     = []int{cost[0], cost[1]}
		result = 0
	)

	for i := 2; i < len(cost); i++ {
		result = min(dp[0], dp[1]) + cost[i]
		dp[0] = dp[1]
		dp[1] = result
	}

	return min(dp[0], dp[1])
}

// 62.不同路径
func UniquePaths(m int, n int) int {
	// 1. 确定 dp[i][j] 下标的含义
	// 到达下标 (i,j) 位置的路径方式
	// 2. 推导 dp[i][j] 表达式
	// dp[i][j] = dp[i-1][j](i>=1) + dp[i][j-1](j>=1)
	// 3. 确定 dp[i][j] 初始值
	// dp[0][0..n-1] = 1
	// dp[0..m-1][0] = 1
	// 4. 遍历顺序
	// 双重遍历(1..m-1)(1..n-1)
	// 5. 举例来确定
	// 输入：m = 2, n = 3 输出：3
	// dp[0][0] = 1
	// dp[0][1] = 1
	// dp[0][2] = 1
	// dp[1][0] = 1
	// dp[1][1] = dp[1][0]] + dp[0][1] = 2
	// dp[1][2] = dp[0][2] + dp[1][1] = 1+2=3

	var (
		dp = make([][]int, m)
	)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	fmt.Println(dp)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

// 63. 有障碍物的 UniquePath
// [[0,0,0],[0,1,0],[0,0,0]]
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		// (1) 确定 dp[i][j] 下标的含义
		// i, j 表示在 obstacleGrid 中的下标
		m  = len(obstacleGrid)
		n  = len(obstacleGrid[0])
		dp = make([][]int, m)
	)

	// 确定 dp[i][j] 初始值
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		// 如果是障碍物, 后面的就都是0, 不用循环了
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {

		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	fmt.Println("我的初始值对了吗》?", dp)

	// (2) 推导 dp 表达式
	// dp[i][j] = dp[i][j-1], [i,j] 不是障碍物时, 左边是障碍物
	// dp[i][j] = dp[i-1][j], [i,j] 不是障碍物时, 上边是障碍物
	// dp[i][j] = 0, [i,j] 就是障碍物

	// (4) 遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// 举例确定

	return dp[m-1][n-1]
}

// 343. 整数拆分
func IntegerBreak(n int) int {
	// 确定 dp 下标含义
	// dp[i] 拆分 i 得到最大的相乘结果
	var dp = make([]int, n+1)

	// 推导 dp 表达式
	// dp[i] = j*(i-j) // 拆分成两个
	// dp[i] = j*dp[i-j] // 拆分成两个以上

	// 确定初始值
	// dp[0], dp[1] 没有意义, 且题目附加条件是 n 不小于 2
	// 但是为了 能够计算, 把 dp[0], dp[1]  都设置为 0
	// dp[2] = 1*1 = 1
	dp[0], dp[1], dp[2] = 0, 0, 1

	// 遍历(1..(i-1))
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	// 验证

	return dp[n]
}
