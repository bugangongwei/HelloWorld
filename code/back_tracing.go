package code

import (
	"fmt"
	"strings"
)

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

// 51. N 皇后
func SolveNQueens(n int) [][]string {
	// 定义变量
	var (
		// 最终结果
		result = make([][]string, 0)
		// 单行结果
		path = make([][]string, n)
		// 递归函数
		// rowIdx 当前正在处理的行号[1...n-1]
		fn func(rowIdx int)
	)

	if n == 0 {
		return result
	}

	for i := 0; i < n; i++ {
		var row = make([]string, n)
		for j := 0; j < n; j++ {
			row[j] = "."
		}
		path[i] = row
	}

	fn = func(rowIdx int) {
		// 终止条件: 已经处理完最后一行
		if rowIdx == n {
			var tmp = make([]string, n)
			for i, row := range path {
				tmp[i] = strings.Join(row, "")
			}
			result = append(result, tmp)
			return
		}

		// 循环遍历: 逐列处理
		for colIdx := 0; colIdx < n; colIdx++ {
			// 处理节点
			if valid(rowIdx, colIdx, path) {
				path[rowIdx][colIdx] = "Q"
				// 递归
				fn(rowIdx + 1)
				// 回溯
				path[rowIdx][colIdx] = "."
			}
		}
	}

	fn(0)
	return result
}

func valid(rowIdx, colIdx int, path [][]string) bool {
	// 第一行的位置都是合法的
	if rowIdx == 0 {
		return true
	}

	// 同列检查
	for k := rowIdx - 1; k >= 0; k-- {
		if path[k][colIdx] == "Q" {
			return false
		}
	}

	// 135 度检查(易出错的点,要用 for 循环)
	for r, c := rowIdx, colIdx; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if path[r][c] == "Q" {
			return false
		}
	}

	// 45 度检查(易出错的点,要用 for 循环)
	for r, c := rowIdx, colIdx; r >= 0 && c < len(path); r, c = r-1, c+1 {
		if path[r][c] == "Q" {
			return false
		}
	}

	return true
}

// N 皇后官方
var res [][]string

func isValid(board [][]string, row, col int) (res bool) {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			if board[0][1] == "Q" && row == 3 && col == 4 {
				fmt.Println("====> 1")
			}
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			if board[0][1] == "Q" && row == 3 && col == 4 {
				fmt.Println("====> 2")
			}
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			if board[0][1] == "Q" && row == 3 && col == 4 {
				fmt.Println("====> 3")
			}
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			if board[0][1] == "Q" && row == 3 && col == 4 {
				fmt.Println("====> 4")
			}
			return false
		}
	}
	return true
}

func backtrack(board [][]string, row int) {
	size := len(board)
	if row == size {
		temp := make([]string, size)
		for i := 0; i < size; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
		return
	}
	for col := 0; col < size; col++ {
		if board[0][1] == "Q" && row == 3 && col == 4 {
			fmt.Println(board, row, col, isValid(board, row, col))
		}

		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrack(board, row+1)
		board[row][col] = "."
	}
}

func SolveNQueensR(n int) [][]string {
	res = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtrack(board, 0)

	return res
}
