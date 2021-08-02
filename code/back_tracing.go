package code

// 77. 组合
/* 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合 */
func Combanation(n, k int) [][]int {
	// 定义变量
	var (
		result [][]int
		path   []int // from root to leaf
	)

	var fn func(i int)
	fn = func(startIndex int) {
		// 终止条件: 达到条件(叶子节点)
		if len(path) == k {
			// 二维数组直接 append path 实际上 append 的是 path 的地址
			var tmp = make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 循环遍历: 遍历本层集合
		restNeeded := k - len(path)
		for i := startIndex; i <= n-restNeeded+1; i++ {
			// 处理节点
			path = append(path, i)
			// 递归
			fn(i + 1)
			// 回溯
			path = path[0 : len(path)-1]
		}
	}

	fn(1)
	return result
}

// 216.组合总和III
func CombinationSum3(k int, n int) [][]int {
	// 定义变量
	var (
		result [][]int
		path   []int // from root to leaf
		sum    int
		m      = 9
	)

	var fn func(i int)
	fn = func(startIndex int) {
		// 剪枝: 超出目标
		if len(path) == k && sum != n {
			return
		}

		// 终止条件: len(path) == k && sum == n
		if len(path) == k && sum == n {
			var tmp = make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 循环遍历
		for i := startIndex; i <= m-(k-len(path))+1; i++ {
			// 处理节点
			path = append(path, i)
			sum += i
			// 递归
			fn(i + 1)
			// 回溯
			sum -= i
			path = path[0 : len(path)-1]
		}
	}

	fn(1)
	return result
}

func LetterCombinations(digits string) []string {
	// 定义变量
	var (
		result    []string
		path      string
		numLetMap = map[string]string{
			"2": "abc",
			"3": "def",
			"4": "ghi",
			"5": "jkl",
			"6": "mno",
			"7": "pqrs",
			"8": "tuv",
			"9": "wxyz",
		}
		fn func(index int)
	)

	if len(digits) == 0 {
		return result
	}

	fn = func(index int) {
		// fmt.Println(index, path, result)
		// 终止条件
		if len(path) == len(digits) {
			result = append(result, path)
			return
		}

		// 循环遍历
		letters := numLetMap[string(digits[index])]
		// fmt.Println(index, letters)
		for i := 0; i < len(letters); i++ {
			// 处理节点
			path += string(letters[i])
			// 递归
			fn(index + 1)
			// 回溯
			path = path[0 : len(path)-1]
		}
	}

	fn(0)
	return result
}
